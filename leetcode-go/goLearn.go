package main

import "fmt"

func main() {
	var test map[int]string
	test = make(map[int]string, 0)
	test[5] = "ll"
	fmt.Println(test)
	ok, z := test[6]
	fmt.Println(z)
	fmt.Println(ok)
}
