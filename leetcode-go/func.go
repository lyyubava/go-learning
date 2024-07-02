package main

import (
	"fmt"
)

func main() {
	b := []int{1, 2, 3}
	c := b
	c[1] = 3
	fmt.Println(c)
}
