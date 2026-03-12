package hashutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSHA256(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "empty", in: "", want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{name: "basic", in: "Goodbye", want: "c015ad6ddaf8bb50689d2d7cbf1539dff6dd84473582a08ed1d15d841f4254f4"},
		{name: "with punctuation", in: "Goodbye my friend!!!", want: "442bc27f58a0ca075b885158cefd8db0ba1053116654e55a4d18fc4504eebc29"},
		{name: "cyrillic", in: "Привет! Как дела?", want: "d758797821233d1d2b4ae7e78a60f3295f274d2c9767ea44b0d04ac9d4787a6e"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := HashSHA256(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}
