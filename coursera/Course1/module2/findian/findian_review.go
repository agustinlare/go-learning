package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a string: ")
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not found!")
	}
}
