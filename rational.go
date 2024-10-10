// Package rational implements rational numbers.
//
// In addition to arithmetic functions Plus, Minus, Times and DivideBy, the package also implements conversion, comparison and conversion from integers.
//
// The package also provides functions for parsing rational numbers from strings and outputting rational numbers as strings.
//
// Rational numbers are value types. Once created, they cannot and should not be changed. The result of all operations is always a new rational number.
package rational

import (
	"fmt"
)

// Rational ist ein Datentyp, der rationale Zahlen repräsentiert.
//
// Rationale Zahlen bestehen aus einem Zähler und einem Nenner.
// Wenn der Nenner den Wert 0 besitzt, ist die rationale Zahl nicht definiert; Sie besitzt dann den Wert NaN (Not a number).
//
// Die Ergebnisse aller Operationen mit undefinierten rationalen Zahl führen wiederum zu einer undefinierten rationalen Zahl.
type Rational struct {
	numerator   int
	denominator int
}

func (r Rational) String() string {
	if r.IsNaN() {
		return "NaN"
	}

	if r.denominator == 1 {
		return fmt.Sprint(r.numerator)
	}

	return fmt.Sprintf("%d/%d", r.numerator, r.denominator)
}

// FromInt converts an integer to a rational number.
func FromInt(numerator int) Rational {
	return Rational{
		numerator:   numerator,
		denominator: 1,
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func (left Rational) Plus(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

	numerator := right.numerator*left.denominator + left.numerator*right.denominator
	denominator := right.denominator * left.denominator
	return Rational{numerator, denominator}.normalize()
}

func (left Rational) Minus(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

	numerator := left.numerator*right.denominator - right.numerator*left.denominator
	denominator := right.denominator * left.denominator
	return Rational{numerator, denominator}.normalize()
}

func (left Rational) Times(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

	numerator := left.numerator * right.numerator
	denominator := left.denominator * right.denominator
	return Rational{numerator, denominator}.normalize()
}

func (left Rational) DivideBy(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

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

func (r Rational) IsNaN() bool {
	return r.denominator == 0
}
