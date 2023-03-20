package main

import (
	"strings"
)

func spinWords(str string) {
	// scan := bufio.NewReader(os.Stdin)
	// input, _ := scan.ReadString('\n')
	// str := "Welcom to CodeWars"
	words := strings.Fields(str)

	for i, v := range words {
		if len(v) >= 5 {
			words[i] = reverseWord(v)
		}
	}
	strings.Join(words, " ")
}

func reverseWord(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
