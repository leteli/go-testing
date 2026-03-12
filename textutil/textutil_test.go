package textutil

import (
	"strings"
	"testing"
)

func assertStringsEqual(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestNormalizeSpaces(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "spaces", in: "  How are      you? ", want: "How are you?"},
		{name: "spaces tabs and newlines", in: " \t\t\t\n\t How\nare\n  \t\t  you? \t ", want: "How are you?"},
		{name: "cyrillic", in: " \n\nКак\nдела?\n", want: "Как дела?"},
		{name: "empty", in: "", want: ""},
		{name: "separators only", in: "    \t\n \n", want: ""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := NormalizeSpaces(tc.in)
			assertStringsEqual(t, got, tc.want)
		})
	}
}

func TestCountLines(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want int
	}{
		{name: "one line", in: "How are you?", want: 1},
		{name: "multiple lines", in: "How\nare\nyou?", want: 3},
		{name: "cyrillic", in: " \n\nКак\nдела?\n", want: 4},
		{name: "empty", in: "", want: 0},
		{name: "newlines only", in: "\n\n\n", want: 3},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			reader := strings.NewReader(tc.in)
			got, err := CountLines(reader)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
