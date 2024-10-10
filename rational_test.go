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
		"1\u22152": {
			number: FromInt(1),
			want:   "1",
		},
		"NaN": {
			number: Rational{},
			want:   "NaN",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := fmt.Sprint(test.number)
			equal(t, got, test.want)
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
		"1\u22152 + 1\u22152 = 1":        {Rational{1, 2}, Rational{1, 2}, FromInt(1)},
		"2 + 3 = 5":                      {FromInt(2), FromInt(3), FromInt(5)},
		"1\u22152 + 1\u22153 = 5\u22156": {Rational{1, 2}, Rational{1, 3}, Rational{5, 6}},
		"NaN + 1\u22152 = NaN":           {Rational{1, 0}, Rational{1, 2}, Rational{}},
		"1\u22152 + NaN = NaN":           {Rational{1, 0}, Rational{1, 2}, Rational{}},
		"NaN + NaN = NaN":                {Rational{}, Rational{}, Rational{}},
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
		"1\u22152 - 1\u22152 = 0":            {Rational{1, 2}, Rational{1, 2}, Rational{0, 1}},
		"2 - 3 = -1":                         {Rational{2, 1}, Rational{3, 1}, Rational{-1, 1}},
		"1\u22153 - 1\u22155 = 2\u221515":    {Rational{1, 3}, Rational{1, 5}, Rational{2, 15}},
		"-1\u22153 - 1\u22155 = -8\u221515":  {Rational{-1, 3}, Rational{1, 5}, Rational{-8, 15}},
		"-1\u22153 - -1\u22155 = -2\u221515": {Rational{-1, 3}, Rational{-1, 5}, Rational{-2, 15}},
		"2\u22153 - 1\u22152 = 1\u22156":     {Rational{2, 3}, Rational{1, 2}, Rational{1, 6}},
		"NaN - 1\u22152 = NaN":               {Rational{1, 0}, Rational{1, 2}, Rational{}},
		"1\u22152 - NaN = Nan":               {Rational{1, 2}, Rational{42, 0}, Rational{}},
		"NaN - NaN = NaN":                    {Rational{5, 0}, Rational{7, 0}, Rational{}},
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
		"1\u22152 * 1\u22152 = 1\u22154":     {Rational{1, 2}, Rational{1, 2}, Rational{1, 4}},
		"2 * 3 = 6":                          {Rational{2, 1}, Rational{3, 1}, Rational{6, 1}},
		"1\u22153 * 1\u22155 = 1\u221515":    {Rational{1, 3}, Rational{1, 5}, Rational{1, 15}},
		"-1\u22153 * -1\u22155 = -1\u221515": {Rational{-1, 3}, Rational{1, 5}, Rational{-1, 15}},
		"1\u22153 * -1\u22155 = -1\u221515":  {Rational{1, 3}, Rational{-1, 5}, Rational{-1, 15}},
		"1\u22152 * 3\u22154 = 3\u22158":     {Rational{1, 2}, Rational{3, 4}, Rational{3, 8}},
		"NaN * 1\u22152 = NaN":               {Rational{}, Rational{1, 2}, Rational{}},
		"1\u22152 * NaN = NaN":               {Rational{1, 2}, Rational{47, 0}, Rational{}},
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
		"8 \u2215 2 = 4":                        {FromInt(8), FromInt(2), FromInt(4)},
		"2\u22153 \u2215 4\u22155 = 5\u22156":   {Rational{2, 3}, Rational{4, 5}, Rational{5, 6}},
		"-2\u22153 \u2215 4\u22155 = -5\u22156": {Rational{-2, 3}, Rational{4, 5}, Rational{-5, 6}},
		"2\u22153 \u2215 -4\u22155 = 5\u22156":  {Rational{2, 3}, Rational{-4, 5}, Rational{-5, 6}},
		"1\u22152 \u2215 3\u22154 = 2\u22153":   {Rational{1, 2}, Rational{3, 4}, Rational{2, 3}},
		"1\u22152 \u2215 NaN = NaN":             {Rational{1, 2}, Rational{}, Rational{}},
		"NaN \u2215 1\u22152 = NaN":             {Rational{}, Rational{1, 2}, Rational{}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.left.DivideBy(test.right)
			equal(t, got, test.want)
		})
	}
}

func TestIsNaN(t *testing.T) {
	tests := map[string]struct {
		r    Rational
		want bool
	}{
		"0\u22150":  {Rational{0, 0}, true},
		"1\u22150":  {Rational{1, 0}, true},
		"-1\u22150": {Rational{-1, 0}, true},
		"0\u22151":  {Rational{0, 1}, false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			equal(t, test.r.IsNaN(), test.want)
		})
	}
}

func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
