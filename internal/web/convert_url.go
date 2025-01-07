package web

import "strings"

// replace whitespaces with underscores
func ReplaceWhitespacesWithUnderscores(url string) string {
	url = strings.ReplaceAll(url, " ", "_")
	return url
}

// replace underscores with whitespaces
func ReplaceUnderscoresWithWhitespaces(url string) string {
	url = strings.ReplaceAll(url, "_", " ")
	return url
}
