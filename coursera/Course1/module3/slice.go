package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sli := make([]int, 0, 3)
	var x int = 1

	for x != 0 {

		fmt.Println("Please enter a Int:")
		scan := bufio.NewReader(os.Stdin)
		line, _ := scan.ReadString('\n')
		line = strings.ToLower(line)
		line = strings.TrimSpace(line)

		if line == "X" || line == "x" {
			x = 0
		} else {

			str, err := strconv.Atoi(line)

			if err != nil {
				panic(err)
			}

			sli = append(sli, str)
			sort.Ints(sli)

			fmt.Println(sli)

		}
	}
}
