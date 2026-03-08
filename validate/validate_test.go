package validate

import (
	"errors"
	"testing"
)

func TestValidateName(t *testing.T) {
	cases := []struct {
		name    string
		v       string
		wantErr error
	}{
		{name: "valid name", v: "Ann", wantErr: nil},
		{name: "empty", v: "", wantErr: ErrEmptyName},
		{name: "spaces only", v: "  \n ", wantErr: ErrEmptyName},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateName(tc.v)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("ValidateName(%q) error %v, wantErr %v", tc.v, err, tc.wantErr)
			}
		})
	}
}
