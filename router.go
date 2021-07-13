package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://amazon.in",
	}
	c := make(chan string)

	for _, val := range links {
		go linkChecker(val, c)
	}

	for i := 0; i <= len(links)-1; i++ {

		fmt.Println(<-c)
	}
}

func linkChecker(v string, c chan string) {
	_, err := http.Get(v)
	if err != nil {
		fmt.Println("Can't Reach: ", v)
		c <- "Can't load"
	} else {
		fmt.Printf("\n%v Is loading Correctly\n", v)
	}
	c <- "loaded Successfully"
}
