package manipulation

import (
	"regexp"
	"strings"
)

func Slugify(text string) string {
	lower := strings.ToLower(text)
	hyphens := strings.ReplaceAll(lower, " ", "-")
	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
	if err != nil {
		return ""
	}
	safe := reg.ReplaceAllString(hyphens, "")
	return strings.Trim(safe, "-")
}

func CountTotalWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func RemoveSuffix(text string, suffix string) string {
	if strings.HasSuffix(text, suffix) {
		return text[:len(text)-1]
	}
	return text
}
