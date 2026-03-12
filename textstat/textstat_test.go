package textstat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want map[string]int
	}{
		{name: "basic", in: "  hello, how are you? fine, and You? fine, thank you! ", want: map[string]int{"and": 1, "are": 1, "fine": 2, "hello": 1, "how": 1, "thank": 1, "you": 3}},
		{name: "with numbers", in: "how old are you? i am 20, and You? i am 30!", want: map[string]int{"20": 1, "30": 1, "am": 2, "and": 1, "are": 1, "how": 1, "i": 2, "old": 1, "you": 2}},
		{name: "punctuation separators", in: "hello!how_are_you???fine,and_You?_fine,_thank.you!", want: map[string]int{"and": 1, "are": 1, "fine": 2, "hello": 1, "how": 1, "thank": 1, "you": 3}},
		{name: "empty", in: "", want: map[string]int{}},
		{name: "separators only", in: "  _ -- -__//. !", want: map[string]int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := WordCount(tc.in)
			assert.Equal(t, tc.want, got)
			assert.NotNil(t, got)
		})
	}
}
