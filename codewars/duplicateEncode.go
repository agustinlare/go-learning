package main

import (
	"strings"
)

func duplicateEncode(word string) int {
	var countNew int
	word = strings.ToLower(word)

	for _, l := range strings.Split(word, "") {
		thisCount := strings.Count(word, l)
		if thisCount > countNew {
			countNew = thisCount
		}
	}
	return countNew
}

// Best Practice
// func DuplicateEncode(word string) (r string) {
// 	word = strings.ToLower(word)

// 	t := map[rune]int{}
// 	for _, c := range word { t[c]++ }

// 	for _, c := range word {
// 	  if t[c] == 1 {
// 		r += "("
// 	  } else {
// 		r += ")"
// 	  }
// 	}

// 	return
//   }
