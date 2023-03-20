package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap somthing
func Swap(arr []int, i int) {
	temp := arr[i]
	arr[i] = arr[i+1]
	arr[i+1] = temp
}

// BubbleSort somthing
func BubbleSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func mainBORRAR() {
	reader := bufio.NewReader(os.Stdin)
	arr := []int{}
	fmt.Print("Enter numbers : ")
	inp, _ := reader.ReadString('\n')
	inp = strings.TrimSuffix(inp, "\n")
	d := strings.Split(inp, " ")
	for i := 0; i < len(d); i++ {
		temp, _ := strconv.Atoi(d[i])
		arr = append(arr, temp)
	}
	BubbleSort(arr)
	fmt.Print("Sorted Array : ")
	fmt.Println(arr)
}
