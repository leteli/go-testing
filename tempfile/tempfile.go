package tempfile

import (
	"os"
	"strings"
)

func WriteLinesToTemp(prefix string, lines []string) (string, error) {
	f, err := os.CreateTemp("", prefix)
	if err != nil {
		return "", err
	}
	defer f.Close()

	content := strings.Join(lines, "\n")
	if _, err := f.WriteString(content); err != nil {
		return "", err
	}
	return f.Name(), nil
}
