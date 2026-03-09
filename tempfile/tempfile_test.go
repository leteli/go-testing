package tempfile

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func assertError(t testing.TB, err error, wantErr bool) {
	t.Helper()
	if err != nil && !wantErr {
		t.Errorf("unexpected error %v", err)
	}
	if err == nil && wantErr {
		t.Fatalf("expected an error, got nil")
	}
}

func assertFatalError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
}

func TestWriteLinesToTemp(t *testing.T) {
	dir := t.TempDir()
	fmt.Println("dir", dir)
	cases := []struct {
		name   string
		prefix string
		lines  []string
		want   string
	}{
		{name: "single line", prefix: "test1", lines: []string{"hello"}, want: "hello"},
		{name: "multiple lines", prefix: "test2", lines: []string{"hello", "hey", "bonjour", "guten tag"}, want: "hello\nhey\nbonjour\nguten tag"},
		{name: "empty", prefix: "test3", lines: []string{}, want: ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path, err := WriteLinesToTemp(tc.prefix, tc.lines)
			assertFatalError(t, err)
			if !strings.Contains(path, tc.prefix) {
				t.Errorf("expected %s to contain %s", path, tc.prefix)
			}
			data, err := os.ReadFile(path)
			assertFatalError(t, err)

			if string(data) != tc.want {
				t.Errorf("expected %s, got %s", tc.want, string(data))
			}
			err = os.Remove(path)
			assertFatalError(t, err)
			if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
				t.Errorf("expected %s to be deleted, but it still exists", path)
			}
		})
	}
}
