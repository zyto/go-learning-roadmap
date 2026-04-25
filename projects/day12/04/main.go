package main

func main() {

}

func findWord(words []string, target string) (int, bool) {
	for i, word := range words {
		if word == target {
			return i, true
		}
	}

	return -1, false
}

func parseLimit(text string) (int, error) {

}
