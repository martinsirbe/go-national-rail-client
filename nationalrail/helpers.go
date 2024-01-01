package nationalrail

import (
	"regexp"
	"strings"
)

// NRCC messages contains HTML tags and unnecessary whitespace.
var re = regexp.MustCompile("<[^>]*>")

func cleanNRCCMessage(input string) string {
	input = strings.ReplaceAll(input, "\n", " ")
	input = strings.Join(strings.Fields(input), " ")
	cleaned := re.ReplaceAllString(input, "")
	return strings.TrimSpace(cleaned)
}
