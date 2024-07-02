package main

import "fmt"

var res []int

func spiralOrder(matrix [][]int) []int {
	res = []int{}
	rowStart := 0
	rowEnd := len(matrix) - 1
	columnStart := 0
	if len(matrix) != 0 {
		columnStart = len(matrix[0]) - 1
	}
	numberOfElements := len(matrix) * (columnStart + 1)
	columnEnd := 0

	for columnEnd <= columnStart && rowEnd >= rowStart {
		for i := columnEnd; i <= columnStart; i++ {
			res = append(res, matrix[rowStart][i])

		}
		if len(res) == numberOfElements {
			break
		}
		rowStart += 1

		for j := rowStart; j <= rowEnd; j++ {

			res = append(res, matrix[j][columnStart])

		}
		if len(res) == numberOfElements {
			break
		}
		columnStart -= 1
		for k := columnStart; k >= columnEnd; k-- {

			res = append(res, matrix[rowEnd][k])
		}
		if len(res) == numberOfElements {
			break
		}
		rowEnd -= 1
		for x := rowEnd; x >= rowStart; x-- {

			res = append(res, matrix[x][columnEnd])
		}
		if len(res) == numberOfElements {
			break
		}

		columnEnd += 1
	}
	return res

}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	r := spiralOrder(matrix)
	fmt.Println(r)
}
