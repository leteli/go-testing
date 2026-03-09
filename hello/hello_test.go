package hello

import (
	"testing"
)

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got error, want bool) {
	t.Helper()
	if got != nil && !want {
		t.Fatalf("unexpected error %v", got)
	}
	if got == nil && want {
		t.Fatalf("expected an error, got nil")
	}
}

func TestHello(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{name: "basic string 1", in: "Go", want: "Hello, Go", wantErr: false},
		{name: "basic string 2", in: "World", want: "Hello, World", wantErr: false},
		{name: "empty", in: "", want: "", wantErr: true},
		{name: "cyrillic", in: "Гофер", want: "Hello, Гофер", wantErr: false},
		{name: "spaces", in: "  Go ", want: "Hello,   Go ", wantErr: false},
		{name: "two words", in: "Anne Marie", want: "Hello, Anne Marie", wantErr: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := Hello(tc.in)
			assertString(t, out, tc.want)
			assertError(t, err, tc.wantErr)
		})
	}
}
