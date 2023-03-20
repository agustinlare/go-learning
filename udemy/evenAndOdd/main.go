package main

import (
	"fmt"
	"strconv"
)

type sli []int

func main() {
	sli := newSlice(10)
	sli.PrintOdd()
}

func newSlice(i int) sli {
	var sli sli
	for i >= 0 {
		sli = append(sli, i)
		i--
	}
	return sli
}

func (s sli) PrintOdd() {
	fmt.Println(s)
	for _, v := range s {
		if v%2 == 0 {
			fmt.Println(strconv.Itoa(v) + " is even")
		} else {
			fmt.Println(strconv.Itoa(v) + " is odd")
		}
	}
}
