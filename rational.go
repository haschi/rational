package rational

import "fmt"

type Rational struct {
	numerator   int
	denominator int
}

func (r Rational) String() string {
	if r.denominator == 1 {
		return fmt.Sprint(r.numerator)
	}

	panic("Not implemented yet")
}

func FromInt(numerator int) Rational {
	return Rational{
		numerator:   numerator,
		denominator: 1,
	}
}
