package main

import (
	"strings"
)

type palabra struct {
	word  string
	count int
}

func high(s string) string {
	var a palabra

	for _, v := range splitWords(strings.TrimSpace(s)) {
		var count int = 0
		for _, l := range strings.Split(v, "") {
			if l != "" {
				count = count + getCharValue(l)
			}
		}

		if count > a.count {
			a.word = v
			a.count = count
		}
	}
	return a.word
}

func splitWords(s string) []string {
	return strings.Fields(s)
}

func getCharValue(s string) int {
	var arr = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var r int

	for i, v := range arr {
		if v == s {
			r = i + 1
		}
	}

	return r
}

// Best Practice
// func wordScore(s string) (scor byte) {
// 	for i := range s {
// 	  scor += s[i] - 96
// 	}
// 	return
//   }

//   func High(s string) string {
// 	var scor, scorNew byte
// 	var word string
// 	for _, wd := range strings.Split(s, " ") {
// 	  scorNew = wordScore(wd)
// 	  if scorNew > scor {
// 		scor = scorNew
// 		word = wd
// 	  }
// 	}
// 	return word
//   }
