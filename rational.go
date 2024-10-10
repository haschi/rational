package rational

import (
	"errors"
	"fmt"
)

type Rational struct {
	numerator   int
	denominator int
}

func (r Rational) String() string {
	if r.denominator == 1 {
		return fmt.Sprint(r.numerator)
	}

	return fmt.Sprintf("%d/%d", r.numerator, r.denominator)
}

func FromInt(numerator int) Rational {
	return Rational{
		numerator:   numerator,
		denominator: 1,
	}
}

func New(numerator, demoniator int) (Rational, error) {
	if demoniator == 0 {
		return Rational{}, errors.New("denominator can not be zero")
	}

	gcd := gcd(numerator, demoniator)
	return Rational{numerator / gcd, demoniator / gcd}, nil
}

func MustRational(a, b int) Rational {
	rational, err := New(a, b)

	if err != nil {
		panic(err)
	}

	return rational
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func (left Rational) Plus(right Rational) Rational {
	numerator := right.numerator*left.denominator + left.numerator*right.denominator
	denominator := right.denominator * left.denominator
	return Rational{numerator, denominator}.normalize()
}

func (left Rational) Minus(right Rational) Rational {
	numerator := left.numerator*right.denominator - right.numerator*left.denominator
	denominator := right.denominator * left.denominator
	return Rational{numerator, denominator}.normalize()
}

func (left Rational) Times(right Rational) Rational {
	numerator := left.numerator * right.numerator
	denominator := left.denominator * right.denominator
	return Rational{numerator, denominator}.normalize()
}

func (left Rational) DivideBy(right Rational) Rational {
	numerator := left.numerator * right.denominator
	denominator := left.denominator * right.numerator
	return Rational{numerator, denominator}.normalize()
}

func (r Rational) normalize() Rational {
	gcd := gcd(r.numerator, r.denominator)
	result := Rational{r.numerator / gcd, r.denominator / gcd}
	if result.denominator < 1 {
		result = Rational{result.numerator * -1, result.denominator * -1}
	}
	return result
}
