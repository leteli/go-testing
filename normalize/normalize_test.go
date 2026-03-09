package normalize

import (
	"testing"
)

func TestClean(t *testing.T) {
	cases := []struct {
		name string
		s    string
		want string
	}{
		{name: "empty", s: "", want: ""},
		{name: "spaces tabs and newlines", s: "    \n\t    \t\t\t\n  ", want: ""},
		{name: "multiple words with spaces", s: "    \n\t  cat \tand\t dog  \t\t\t\n  ", want: "cat and dog"},
		{name: "case", s: "  CrOcOdiLE ", want: "crocodile"},
		{name: "normalized", s: "monkey", want: "monkey"},
		{name: "punct", s: "  Hello! # How r u ? # I'm fine ! #", want: "hello! # how r u ? # i'm fine ! #"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := Clean(tc.s)
			if out != tc.want {
				t.Errorf("Clean(%s) = %s, got %v", tc.s, tc.want, out)
			}
		})
	}
}
