package words

import "strings"

func NormalizeAll(items []string) []string {
	items = filterNonEmpty(items)
	dst := make([]string, 0, len(items))

	for _, item := range items {
		dst = append(dst, normalize(item))
	}

	return dst
}

func normalize(word string) string {
	return strings.ToLower(clean(word))
}

func clean(word string) string {
	return strings.TrimSpace(word)
}
