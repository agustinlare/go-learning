package main

import (
	"strings"
)

func duplicateCount(s1 string) int {
	var countNew int
	var lastl string
	s1 = strings.ToLower(s1)

	for _, l := range strings.Split(s1, "") {
		thisCount := strings.Count(s1, l)

		if thisCount > 1 && !strings.Contains(lastl, l) {
			lastl = lastl + l
			countNew++
		}
	}
	return countNew
}

// Best Practice
// func duplicate_count(s string) (c int) {
//     h := map[rune]int{}
//     for _, r := range strings.ToLower(s) {
//       if h[r]++; h[r] == 2 { c++ }
//     }
//     return
// }
