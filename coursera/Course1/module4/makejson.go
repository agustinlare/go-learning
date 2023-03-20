package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
1. Prompt user to enter name, and then addres
2. Create map and append the two before using the keys “name” and “address”
3. Use Marshal() to create json
4. Print json
*/

func insertUser(msg string) string {
	fmt.Printf("Insert " + msg + ":")
	scan := bufio.NewReader(os.Stdin)
	str, err := scan.ReadString('\n')
	str = strings.TrimSpace(str)

	handleErr(err)
	return str
}

func handleErr(msg error) {
	if msg != nil {
		panic(msg)
	}
	return
}

func mainBORRAR() {
	var x int = 0

	for x != 1 {
		var name string = "name"
		nameV := insertUser(name)

		if len(nameV) > 0 {
			var addr string = "address"
			addrV := insertUser("address")

			if len(addrV) > 0 {
				myMap := make(map[string]string)

				myMap[addr] = addrV
				myMap[name] = nameV
				j, err := json.Marshal(myMap)

				handleErr(err)

				fmt.Println(string(j))

				x = 1

			} else {
				fmt.Printf("Try again, the address was empty")
			}
		} else {
			fmt.Printf("Cannot be empty")
		}
	}
}
