package main

import (
	"fmt"
	"errors"
)



func maxSumOfK(arr []int, k int) (int, error) {
	n := len(arr)
	if k < n {
		return -1, errors.New("Invalid input: k cannot be bigger than n")
	}
	window := arr[:k]

	for true {
		
	}
}
