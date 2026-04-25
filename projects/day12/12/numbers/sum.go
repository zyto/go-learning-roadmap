package numbers

import "fmt"

func SumFromStrings(items []string) (int, error) {
	sum := 0
	for i, item := range items {
		val, err := ParsePositive(item)
		if err != nil {
			return 0, fmt.Errorf("item %d: %w", i, err)
		}
		sum += val
	}
	return sum, nil
}
