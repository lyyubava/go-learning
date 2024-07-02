package main

import "fmt"

func test(sli [3]int) {
	sli[0] = sli[0] + 1
}

func main() {
	a := [3]int{1, 2, 3}
	test(a)
	fmt.Println(a)
}
