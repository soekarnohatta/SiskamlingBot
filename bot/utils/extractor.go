package utils

import "strings"

func ExtractBool(text string) bool {
	return strings.ToLower(text) == "true"
}
