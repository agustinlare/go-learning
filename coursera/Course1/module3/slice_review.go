package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func insert(s []int, x int) []int {
	i := sort.SearchInts(s, x)
	if i == len(s) {
		return append(s, x)
	} else {
		s = append(s, 0)
		copy(s[i+1:], s[i:])
		s[i] = x
		return s
	}
}

func main1() {
	s := make([]int, 0, 3)
	var in string
	var x int
	for {
		//read a string
		fmt.Print("Enter an integer: ")
		_, err := fmt.Scan(&in)
		if err != nil {
			fmt.Println("Scan error...")
			continue
		}

		// exit on "x"
		if strings.ToLower(in) == "x" {
			os.Exit(0)
		}

		// convert to int
		x, err = strconv.Atoi(in)
		if err != nil {
			fmt.Println("Not an integer...")
			continue
		}

		// insert and print
		s = insert(s, x)
		printSlice(s)
	}
}
