package examples

import "strings"

func lettersCount(s string) map[string]int {
	splitted := strings.Split(s, "")
	countMap := map[string]int{}

	for _, char := range splitted {
		entry, exists := countMap[char]
		if exists {
			countMap[char] = entry + 1
			continue
		}

		countMap[char] = 1
	}

	return countMap
}
