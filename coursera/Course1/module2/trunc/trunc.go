package main

import "fmt"

func maina() {
	var input float64

	fmt.Printf("Insert some float:\n")
	num, err := fmt.Scan(&input)
	var output int = int(input)

	if err != nil || num == 0 {
		panic(fmt.Sprintf("%v", err))
	} else {
		fmt.Println(output)
	}
}
