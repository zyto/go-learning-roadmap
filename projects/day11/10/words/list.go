package words

func NormalizeAll(items []string) []string {
	items = FilterNonEmpty(items)
	dst := make([]string, 0, len(items))

	for _, item := range items {
		dst = append(dst, Normalize(item))
	}

	return dst
}

func CountUnique(items []string) int {
	items = NormalizeAll(items)
	m := make(map[string]bool)
	dst := make([]string, 0, len(items))

	for _, item := range items {
		if m[item] == false {
			m[item] = true
			dst = append(dst, item)
		}
	}

	return len(dst)
}

func FilterNonEmpty(items []string) []string {
	dst := make([]string, 0, len(items))

	for _, item := range items {
		if item != "" {
			dst = append(dst, item)
		}
	}

	return dst
}
