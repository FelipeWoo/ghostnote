package cli

import (
	"fmt"
	"regexp"
	"strings"
)

func normalize(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("input too short")
	}
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	re := regexp.MustCompile(`[^a-zA-Z0-9_\-]`)
	input = re.ReplaceAllString(input, "_")

	return input, nil
}

func cleanCSV(input string) []string {
	if strings.TrimSpace(input) == "" {
		return []string{}
	}

	parts := strings.Split(input, ",")
	out := make([]string, 0, len(parts))
	seen := make(map[string]struct{})

	for _, p := range parts {
		p, _ = normalize(p)

		if p == "" {
			continue
		}

		// avoid duplicates
		if _, exists := seen[p]; exists {
			continue
		}

		seen[p] = struct{}{}
		out = append(out, p)
	}

	return out
}

func cleanPayload(input string) string {
	output := strings.TrimSpace(input)
	output = strings.ReplaceAll(output, `\n`, "\n")
	return output
}
