package utils

import (
	"regexp"
	"strings"
)

func CleanCode(raw string) string {
	re := regexp.MustCompile("```(?:\\w+)?\\s*([\\s\\S]*?)\\s*```")
	matches := re.FindStringSubmatch(raw)

	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}

	return strings.TrimSpace(raw)
}