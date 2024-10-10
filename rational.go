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
	Numerator   int
	Denominator int
}

func (r Rational) String() string {
	if r.IsNaN() {
		return "NaN"
	}

	if r.Denominator == 1 {
		return fmt.Sprint(r.Numerator)
	}

	return fmt.Sprintf("%d/%d", r.Numerator, r.Denominator)
}

// FromInt converts an integer to a rational number.
func FromInt(numerator int) Rational {
	return Rational{
		Numerator:   numerator,
		Denominator: 1,
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

	numerator := right.Numerator*left.Denominator + left.Numerator*right.Denominator
	denominator := right.Denominator * left.Denominator
	return Rational{numerator, denominator}.normalize()
}

func (left Rational) Minus(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

	numerator := left.Numerator*right.Denominator - right.Numerator*left.Denominator
	denominator := right.Denominator * left.Denominator
	return Rational{numerator, denominator}.normalize()
}

func (left Rational) Times(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

	numerator := left.Numerator * right.Numerator
	denominator := left.Denominator * right.Denominator
	return Rational{numerator, denominator}.normalize()
}

func (left Rational) DivideBy(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

	numerator := left.Numerator * right.Denominator
	denominator := left.Denominator * right.Numerator
	return Rational{numerator, denominator}.normalize()
}

func (r Rational) normalize() Rational {
	gcd := gcd(r.Numerator, r.Denominator)
	result := Rational{r.Numerator / gcd, r.Denominator / gcd}
	if result.Denominator < 1 {
		result = Rational{result.Numerator * -1, result.Denominator * -1}
	}
	return result
}

func (r Rational) IsNaN() bool {
	return r.Denominator == 0
}
