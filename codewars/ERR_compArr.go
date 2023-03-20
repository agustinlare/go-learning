package main

import (
	"math"
)

// func main() {
// 	var a []int
// 	var b []int
// 	a = []int{121, 144, 19, 161, 19, 144, 19, 11}
// 	// b = []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
// 	b = []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
// 	// b = nil

// 	fmt.Println(comp(a, b))
// }

// https://www.codewars.com/kata/550498447451fbbd7600041c/train/go

func comp(array1 []int, array2 []int) bool {

	if len(array1) != len(array2) || array1 == nil || array2 == nil {
		return false
	}

	for _, v := range array2 {
		f := math.Sqrt(float64(v))
		if !contains(array1, f) {
			return false
		}
	}
	return true
}

func contains(ar []int, e float64) bool {
	for _, a := range ar {
		if float64(a) == e {
			return true
		}
	}
	return false
}
