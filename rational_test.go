package rational

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tests := map[string]struct {
		number Rational
		want   string
	}{
		"Zero": {
			number: FromInt(0),
			want:   "0",
		},
		"One": {
			number: FromInt(42),
			want:   "42",
		},
		"1/2": {
			number: FromInt(1),
			want:   "1",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := fmt.Sprint(test.number)
			if got != test.want {
				t.Errorf("want %v; got %v", test.want, got)
			}
		})
	}
}

func TestFromInt(t *testing.T) {
	r1 := FromInt(0)
	r2 := FromInt(0)

	if r1 != r2 {
		t.Errorf("want %[1]v == %[2]v; got %[1]v != %[2]v", r1, r2)
	}
}

func TestGcd(t *testing.T) {
	tests := map[string]struct {
		a    int
		b    int
		want int
	}{
		"0 gcd 1":   {0, 1, 1},
		"1 gcd 0":   {1, 0, 1},
		"2 gcd 1":   {2, 1, 1},
		"-5 gcd 10": {5, 10, 5},
		"42 gcd 11": {42, 22, 2},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := gcd(test.a, test.b)
			if got != test.want {
				t.Errorf("got %v; want %v", got, test.want)
			}
		})
	}
}

func TestPlus(t *testing.T) {
	tests := map[string]struct {
		left, right, want Rational
	}{
		"1/2 + 1/2 = 1":   {Rational{1, 2}, Rational{1, 2}, FromInt(1)},
		"2 + 3 = 5":       {FromInt(2), FromInt(3), FromInt(5)},
		"1/2 + 1/3 = 5/6": {Rational{1, 2}, Rational{1, 3}, Rational{5, 6}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.left.Plus(test.right)
			equal(t, got, test.want)
		})
	}
}

func TestMinus(t *testing.T) {
	tests := map[string]struct {
		left, right, want Rational
	}{
		"1/2 - 1/2 = 0":       {Rational{1, 2}, Rational{1, 2}, Rational{0, 1}},
		"2 - 3 = -1":          {Rational{2, 1}, Rational{3, 1}, Rational{-1, 1}},
		"1/3 - 1/5 = 2/15":    {Rational{1, 3}, Rational{1, 5}, Rational{2, 15}},
		"-1/3 - 1/5 = -8/15":  {Rational{-1, 3}, Rational{1, 5}, Rational{-8, 15}},
		"-1/3 - -1/5 = -2/15": {Rational{-1, 3}, Rational{-1, 5}, Rational{-2, 15}},
		"2/3 - 1/2 = 1/6":     {Rational{2, 3}, Rational{1, 2}, Rational{1, 6}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.left.Minus(test.right)
			equal(t, got, test.want)
		})
	}
}

func TestTimes(t *testing.T) {
	tests := map[string]struct {
		left, right, want Rational
	}{
		"1/2 * 1/2 = 1/4":     {Rational{1, 2}, Rational{1, 2}, Rational{1, 4}},
		"2 * 3 = 6":           {Rational{2, 1}, Rational{3, 1}, Rational{6, 1}},
		"1/3 * 1/5 = 1/15":    {Rational{1, 3}, Rational{1, 5}, Rational{1, 15}},
		"-1/3 * -1/5 = -1/15": {Rational{-1, 3}, Rational{1, 5}, Rational{-1, 15}},
		"1/3 * -1/5 = -1/15":  {Rational{1, 3}, Rational{-1, 5}, Rational{-1, 15}},
		"1/2 * 3/4 = 3/8":     {Rational{1, 2}, Rational{3, 4}, Rational{3, 8}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.left.Times(test.right)
			equal(t, got, test.want)
		})
	}
}

func TestDivideBy(t *testing.T) {
	tests := map[string]struct {
		left, right, want Rational
	}{
		"8 / 2 = 4":         {FromInt(8), FromInt(2), FromInt(4)},
		"2/3 / 4/5 = 5/6":   {Rational{2, 3}, Rational{4, 5}, Rational{5, 6}},
		"-2/3 / 4/5 = -5/6": {Rational{-2, 3}, Rational{4, 5}, Rational{-5, 6}},
		"2/3 / -4/5 = 5/6":  {Rational{2, 3}, Rational{-4, 5}, Rational{-5, 6}},
		"1/2 / 3/4 = 2/3":   {Rational{1, 2}, Rational{3, 4}, Rational{2, 3}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.left.DivideBy(test.right)
			equal(t, got, test.want)
		})
	}
}

func equal[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
