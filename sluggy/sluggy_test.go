package sluggy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlug(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "space separator", in: "test function 1", want: "test-function-1"},
		{name: "punctuation ", in: ".test.function. ; 2", want: "test-function-2"},
		{name: "repeated separators", in: " __ test____function__ ", want: "test-function"},
		{name: "different case", in: "   tESt    fUNCtiON ", want: "test-function"},
		{name: "cyrillic uppercased", in: "ТЕСТОВАЯ   ФУНКЦИЯ ", want: "тестовая-функция"},
		{name: "empty string", in: "", want: ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Slug(tc.in)
			assert.Equal(t, tc.want, got, "Slug(%s)", tc.in)
		})
	}
}
