package even

import "testing"

func TestIsEven(t *testing.T) {
	testCases := []struct {
		name string
		v    int
		want bool
	}{
		{name: "even", v: 10, want: true},
		{name: "odd", v: 17, want: false},
		{name: "zero", v: 0, want: true},
		{name: "negative_even", v: -988, want: true},
		{name: "negative_odd", v: -111, want: false},
		{name: "large", v: 100_000, want: true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsEven(tc.v)
			if got != tc.want {
				t.Errorf("IsEven(%d) = %v; want %v", tc.v, got, tc.want)
			}
		})
	}
}
