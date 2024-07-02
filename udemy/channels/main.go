package main

import (
	"fmt"
	"net/http"
	"time"
)

func evaluate(a int) func(f func()) func() {
	f = func() {
		res := a == 1 || a == 2 || a == 3
	}
	return func()
}

func main() {
	var a int
	a = 2
	b := (a == 2)
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://laaaamazon1.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)

	}
	for l := range c {

		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)

	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // blocking call
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}
	fmt.Println(link, " is up!")
	c <- link

}
