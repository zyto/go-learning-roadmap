package words

import "strings"

func Normalize(word string) string {
	return strings.ToLower(clean(word))
}

func clean(word string) string {
	return strings.TrimSpace(word)
}
