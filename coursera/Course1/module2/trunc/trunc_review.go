package main

import (
	"fmt"
	"math"
)

func asd() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println("After truncated: ", int(math.Round(input)))
}
