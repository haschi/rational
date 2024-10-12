// Package rational implements rational numbers.
//
// In addition to arithmetic functions Plus, Minus, Times and DivideBy, the package also implements conversion and comparison.
//
// Rational numbers are value types. Once created, they cannot and should not be changed. The result of all operations are always a new rational number.
package rational

import (
	"fmt"
)

// Rational is a data type that represents rational numbers.
//
// Rational numbers consist of a numerator and a denominator.
//
// If the denominator has the value 0, the rational number is not defined.
// They are then not a number (NaN).
// This can be checked using the IsNaN method.
// Operations with operands for which IsNaN is true in turn produce a result for which IsNaN is also true.
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

// Plus calculates the sum of left and right.
//
// The result is normalized.
//
// If IsNaN is true for one of the two operands, the method returns a rational number for which IsNan is also true.
func (left Rational) Plus(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

	return Rational{
		Numerator:   right.Numerator*left.Denominator + left.Numerator*right.Denominator,
		Denominator: right.Denominator * left.Denominator,
	}.normalize()
}

// Minus calculates the difference between left and right.
//
// The result is normalized.
//
// If IsNaN is true for one of the two operands, the method returns a rational number for which IsNan is also true.
func (left Rational) Minus(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

	return Rational{
		Numerator:   left.Numerator*right.Denominator - right.Numerator*left.Denominator,
		Denominator: right.Denominator * left.Denominator,
	}.normalize()
}

// Times calculates the product of left and right.
//
// The result is normalized.
//
// If IsNaN is true for one of the two operands, the method returns a rational number for which IsNan is also true.
func (left Rational) Times(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

	numerator := left.Numerator * right.Numerator
	denominator := left.Denominator * right.Denominator
	return Rational{numerator, denominator}.normalize()
}

// DivideBy calculates the quotient of left and right.
//
// The result is normalized.
//
// If IsNaN is true for one of the two operands, the method returns a rational number for which IsNan is also true.
func (left Rational) DivideBy(right Rational) Rational {
	if left.IsNaN() || right.IsNaN() {
		return Rational{}
	}

	numerator := left.Numerator * right.Denominator
	denominator := left.Denominator * right.Numerator
	return Rational{numerator, denominator}.normalize()
}

// IsNaN reports whether the value r is a valid rational number.
func (r Rational) IsNaN() bool {
	return r.Denominator == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func (r Rational) normalize() Rational {
	gcd := gcd(r.Numerator, r.Denominator)
	result := Rational{r.Numerator / gcd, r.Denominator / gcd}
	if result.Denominator < 1 {
		result = Rational{result.Numerator * -1, result.Denominator * -1}
	}
	return result
}
