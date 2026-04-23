package words

import "strings"

func Normalize(word string) string {
	return strings.ToLower(strings.TrimSpace(word))
}
