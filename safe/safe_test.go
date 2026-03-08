package safe

import "testing"

func TestMustAt(t *testing.T) {
	in := []int{1, 2, 5, 8, 9}
	cases := []struct {
		name  string
		v     int
		exp   int
		panic bool
	}{
		{name: "valid index", v: 2, exp: in[2], panic: false},
		{name: "zero index", v: 0, exp: in[0], panic: false},
		{name: "last index", v: len(in) - 1, exp: in[len(in)-1], panic: false},
		{name: "index greater than len", v: 8, exp: 0, panic: true},
		{name: "negative index", v: -8, exp: 0, panic: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if tc.panic {
					if r == nil {
						t.Errorf("expected panic %q, got none", ErrOutOfRange)
					} else if r != ErrOutOfRange {
						t.Errorf("expected panic %q, got %v", ErrOutOfRange, r)
					}
				} else {
					if r != nil {
						t.Fatalf("unexpected panic %v", r)
					}
				}
			}()
			out := MustAt(in, tc.v)
			if out != tc.exp {
				t.Errorf("MustAt(%v, %d) expected %v received %v", in, tc.v, tc.exp, out)
			}
		})
	}
}
