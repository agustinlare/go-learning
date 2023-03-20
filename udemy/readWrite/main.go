package main

import (
	"io"
	"os"
)

func main() {

	file, err := os.Open(os.Args[1])

	io.Copy(os.Stdout, file)

	if err != nil {
		panic(err)
	}

	defer file.Close()
}
