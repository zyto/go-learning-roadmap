package words

func filterNonEmpty(items []string) []string {
	dst := make([]string, 0, len(items))

	for _, item := range items {
		if item != "" {
			dst = append(dst, item)
		}
	}

	return dst
}
