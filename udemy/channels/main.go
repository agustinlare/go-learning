package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}
	c := make(chan string)

	for _, v := range sites {
		go httpGet(v, c)
	}

	// Even if the channel returns a string
	// also can be used with the range keyword.
	// We should never try to access the same variable from another child routine.
	// We pass l to the function literal so the value of the variable don't get
	// change on another routine.
	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			httpGet(l, c)
		}(l)
	}
}

func httpGet(s string, c chan string) {
	_, err := http.Get(s)

	if err != nil {
		fmt.Println(s, "might be down!")
		c <- s
	} else {
		fmt.Println(s, "it's up!")
		c <- s
	}

	return
}
