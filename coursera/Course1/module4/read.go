package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func isErr(msg error) {
	if msg != nil {
		panic(msg)
	}
	return
}

type person struct {
	fname string
	lname string
}

func newPerson(name string, lastname string) *person {
	var per person
	per.fname = name
	per.lname = lastname
	return &per
}

func main() {
	fmt.Printf("Please input the file path or file name:")
	scan := bufio.NewReader(os.Stdin)
	fileName, err := scan.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	isErr(err)

	//file, err := os.Open("golang/module4/file.txt")
	file, err := os.Open(fileName)
	isErr(err)
	defer file.Close()

	rd := bufio.NewReader(file)

	sli := make([]person, 0)

	for {
		// bline, _, err := rd.ReadLine()
		bline, err := rd.ReadString('\n')
		newLine := strings.TrimSpace(string(bline))
		nl := strings.Fields(newLine)

		p1 := new(person)
		p1.fname = nl[0]
		p1.lname = nl[1]

		sli = append(sli, *p1)

		if err == io.EOF {
			break
		}
	}

	for a, b := range sli {
		fmt.Println(a, b)
	}
}
