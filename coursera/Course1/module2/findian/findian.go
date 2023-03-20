package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainasd() {
	fmt.Println("Please enter a string:")

	scan := bufio.NewReader(os.Stdin)
	line, _ := scan.ReadString('\n')

	newStr := strings.ToLower(line)
	newStr = strings.TrimSpace(newStr)

	if strings.HasPrefix(newStr, "i") && strings.Contains(newStr, "a") && strings.HasSuffix(newStr, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found!")
	}

}
