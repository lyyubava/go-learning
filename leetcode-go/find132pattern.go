package main

import (
	"fmt"
	"math"
)

func insertInStack(s *[]int, el int) {
	*s = append(*s, el)
}

func popStack(s *[]int) int {
	poped := (*s)[len((*s))-1]
	(*s) = (*s)[:len((*s))-1]
	return poped
}

func isEmpty(s *[]int) bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func getStackTop(s *[]int) int {
	return (*s)[len((*s))-1]
}

func find132pattern(nums []int) bool {
	c := math.MinInt32
	s := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < c {
			return true
		}
		for !isEmpty(&s) && nums[i] > getStackTop(&s) {
			c = popStack(&s)
		}
		insertInStack(&s, nums[i])

	}
	return false
}
func main() {
	nums := []int{3, 5, 0, 3, 4}
	res := find132pattern(nums)
	fmt.Println(res)
}
