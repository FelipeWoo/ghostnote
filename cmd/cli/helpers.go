package cli

import (
	"fmt"
	"regexp"
	"strings"
)

func normalize(input string) (string, error) {
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	re := regexp.MustCompile(`[^a-zA-Z0-9_\-]`)
	input = re.ReplaceAllString(input, "_")

	if len(input) < 3 {
		return "", fmt.Errorf("title too short")
	}
	return input, nil
}
