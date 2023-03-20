package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	inputSlice := make([]int, 0, 3)
	var input string
	fmt.Printf("insert floating point number: ")
	for input != "X" {
		fmt.Scan(&input)
		inputInt, err := strconv.Atoi(input)
		if err != nil && input != "X" {
			fmt.Printf("must be and integer: ")
			continue
		}
		if input != "X" {
			inputSlice = append(inputSlice, inputInt)
			sort.Ints(inputSlice)
			fmt.Println(inputSlice)
			fmt.Printf("insert next point number: ")
		} else {
			fmt.Printf("final slice: %d", inputSlice)
		}
	}
}
