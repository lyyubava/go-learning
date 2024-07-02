package main

import (
	"fmt"
	"math"
	"strconv"
)

func validateArr(arr []byte) bool {
	elements := make(map[string]bool, 0)
	for _, el := range arr {
		if string(el) == "." {
			continue
		} else {
			num, _ := strconv.Atoi(string(el))

			if num < 1 || num > 9 {
				return false
			}
			_, ok := elements[string(el)]
			if !ok {
				elements[string(el)] = true
			} else {
				return false
			}

		}
	}
	return true

}
func transposeMatrix(matrix [][]byte) [][]byte {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func isValidSudoku(board [][]byte) bool {

	for _, el := range board {
		validateHorizontal := validateArr(el)
		if !validateHorizontal {
			return false
		}
	}
	for _, el := range transposeMatrix(board) {
		validateVertical := validateArr(el)
		if !validateVertical {
			return false
		}
	}
	tmp := make([]byte, 0)
	N := int(math.Sqrt(float64(len(board))))
	for i := 0; i < len(board); i = i + N {
		j := 0
		k := N
		for j < len(board[i]) {
			if k == N {
				validateTmp := validateArr(tmp)
				if !validateTmp {

					return false
				}
				tmp = make([]byte, 0)
				k = 0
				continue
			}
			tmp = append(tmp, board[j][i:i+N]...)
			j++
			k++
		}
		validateTmp := validateArr(tmp)
		if !validateTmp {

			return false
		}

	}
	return true

}

func main() {
	arr := make([][]byte, 0)

	arr = append(arr, []byte("......5.."))
	arr = append(arr, []byte("........."))
	arr = append(arr, []byte("........3"))
	arr = append(arr, []byte(".2.7....."))
	arr = append(arr, []byte("8365....1"))
	arr = append(arr, []byte(".....1..."))
	arr = append(arr, []byte("2.......5"))
	arr = append(arr, []byte("........7"))
	arr = append(arr, []byte("...4...7."))

	res := isValidSudoku(arr)
	fmt.Println(res)

}
