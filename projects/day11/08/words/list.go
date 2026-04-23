package words

func NormalizeAll(items []string) []string {
	dst := make([]string, 0, len(items))

	for _, item := range items {
		dst = append(dst, clean(item))
	}

	return dst
}
