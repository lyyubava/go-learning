package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	for _, el := range a {
		go func() {
			printer(el)
		}()
	}
}

func printer(i int) {
	fmt.Println(i)
}
