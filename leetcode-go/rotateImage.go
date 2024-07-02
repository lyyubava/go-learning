package main

import "fmt"

func transposeMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		k := i + 1
		for j := k; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

		}
	}
}
func reverseSlice(s []int) {
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func rotate(matrix [][]int) {
	transposeMatrix(matrix)
	for i := 0; i < len(matrix); i++ {
		reverseSlice(matrix[i])
	}
}
func main() {
	matrix := make([][]int, 0)
	matrix = append(matrix, []int{1, 2, 3})
	matrix = append(matrix, []int{4, 5, 6})
	matrix = append(matrix, []int{7, 8, 9})
	rotate(matrix)
	fmt.Println(matrix)

}
