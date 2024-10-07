package rational_test

import (
	"fmt"
	"testing"

	"github.com/haschi/rational"
)

func TestString(t *testing.T) {
	tests := map[string]struct {
		number rational.Rational
		want   string
	}{
		"Zero": {
			number: rational.FromInt(0),
			want:   "0",
		},
		"One": {
			number: rational.FromInt(42),
			want:   "42",
		},
		"1/2": {
			number: rational.FromInt(1),
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
	r1 := rational.FromInt(0)
	r2 := rational.FromInt(0)

	if r1 != r2 {
		t.Errorf("want %v == %v; got %v != %v", r1, r2, r1, r2)
	}
}
