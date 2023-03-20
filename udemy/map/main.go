package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"

	delete(colors, "white")

	// fmt.Printf("%+v", colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for i, v := range c {
		fmt.Println(i, v)
	}
}
