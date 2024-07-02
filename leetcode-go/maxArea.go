package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxAreaBruteForce(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			if max == 0 || max < (j-i)*min(height[j], height[i]) {
				max = (j - i) * min(height[j], height[i])
			}

		}

	}
	return max
}

func maxArea(height []int) int {
	i, res, j := 0, 0, len(height)-1
	area := 0

	for i < j {
		if height[i] < height[j] {
			res = height[i] * (j - i)
			i++
		} else {
			res = height[j] * (j - i)
			j--
		}
		if res > area {
			area = res
		}
	}
	return area
}

func main() {
	var res int
	height := []int{1, 1}
	res = maxArea(height)
	fmt.Println(res)
}
