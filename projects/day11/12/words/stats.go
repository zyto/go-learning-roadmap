package words

func CountUnique(items []string) int {
	m := make(map[string]bool)
	unique := 0

	for _, item := range items {
		if m[item] == false {
			m[item] = true
			unique++
		}
	}

	return unique
}
