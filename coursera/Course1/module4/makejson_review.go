package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func mainBORRAR2() {
	inputr := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter Name: ")
	name, _ := inputr.ReadString('\n')

	fmt.Printf("Enter Address: ")
	address, _ := inputr.ReadString('\n')

	naMap := make(map[string]string)
	naMap["name"] = strings.Trim(name, "\n")
	naMap["address"] = strings.Trim(address, "\n")

	jsonString, err := json.Marshal(naMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonString))

}
