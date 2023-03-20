package main

func getCount(s string) int {
	// s := "Hola como estas agustin lavarello"
	var count int
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, c := range s {
		for _, v := range vowels {
			if v == c {
				count = count + 1
			}
		}
	}

	return count
}
