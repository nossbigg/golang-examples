package examples

import "strings"

type lettersCountMap map[string]int

func lettersCount(s string) lettersCountMap {
	splitted := strings.Split(s, "")
	countMap := lettersCountMap{}

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
