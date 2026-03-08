package validate

import (
	"errors"
	"strings"
)

var ErrEmptyName = errors.New("empty name")

func ValidateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrEmptyName
	}
	return nil
}
