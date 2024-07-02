package main

import (
	"fmt"
	"math"
)

var sudokuValues []byte = []byte("123456789")

// [[".",".","9","7","4","8",".",".","."],
// ["7",".",".",".",".",".",".",".","."],
// [".","2",".","1",".","9",".",".","."],
// [".",".","7",".",".",".","2","4","."],
// [".","6","4",".","1",".","5","9","."],
// [".","9","8",".",".",".","3",".","."],
// [".",".",".","8",".","3",".","2","."],
// [".",".",".",".",".",".",".",".","6"],
// [".",".",".","2","7","5","9",".","."]]

func main() {
	arr := make([][]byte, 0)

	arr = append(arr, []byte("..9748..."))
	arr = append(arr, []byte("7........"))
	arr = append(arr, []byte(".2.1.9..."))
	arr = append(arr, []byte("..7...24."))
	arr = append(arr, []byte(".64.1.59."))
	arr = append(arr, []byte(".98...3.."))
	arr = append(arr, []byte("...8.3.2."))
	arr = append(arr, []byte("........6"))
	arr = append(arr, []byte("...2759.."))

	//indexes := getAllIndexesForElementsInSquare([2]int{4, 7}, int(math.Sqrt(float64(9))))
	solveSudoku(arr)
	decisions := make([]map[[2]int]byte, 0)
	makeNewDecision(arr, decisions)
	fmt.Println(decisions)
	//fmt.Println(getPossibleValues(arr, [2]{}))
}

func findNumberOfUnsolved(board [][]byte) int {
	unsolved := 0
	for i := 0; i < 0; i++ {
		for j := 0; j < 0; j++ {
			if string(board[i][j]) != "." {
				continue
			}
			unsolved += 1
		}
	}
	return unsolved
}
func compareTwoBoards(board1 [][]byte, board2 [][]byte) bool {
	return findNumberOfUnsolved(board1) == findNumberOfUnsolved(board2)

}

// builds map for unknown elements in the following structure:
//
//		{
//		 "index" : [ 1, 3, 4],
//	    }
func BuildMapFromBoardForPossibleValues(board [][]byte) map[[2]int][]byte {
	m := make(map[[2]int][]byte)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if string(board[i][j]) != "." {
				continue
			}
			m[[2]int{i, j}] = getPossibleValues(board, [2]int{i, j})
		}
	}
	return m
}
func getAllIndexesForElementsInSquare(index [2]int, N int) [][2]int {
	var squareIndexes [][2]int
	upperBoundRow := N*(index[0]/N) + N
	lowerBoundRow := N * (index[0] / N)
	upperBoundColumn := N*(index[1]/N) + N
	lowerBoundColumn := N * (index[1] / N)
	for i := lowerBoundRow; i < upperBoundRow; i++ {
		for j := lowerBoundColumn; j < upperBoundColumn; j++ {
			squareIndexes = append(squareIndexes, [2]int{i, j})
		}

	}
	return squareIndexes
}
func determineInIntervals(intervals map[[2]int][]byte) (byte, [2]int, bool) {
	type indexesNumOfOccureneces struct {
		numOfOccureneces int
		indexes          [][2]int
	}
	occuranceMap := make(map[byte]indexesNumOfOccureneces)
	for index, interval := range intervals {
		for _, i := range interval {
			if _, ok := occuranceMap[i]; ok {
				ind := occuranceMap[i].indexes
				ind = append(ind, index)
				occuranceMap[i] = indexesNumOfOccureneces{numOfOccureneces: occuranceMap[i].numOfOccureneces + 1, indexes: ind}

			} else {
				ind := [][2]int{}
				ind = append(ind, index)
				occuranceMap[i] = indexesNumOfOccureneces{numOfOccureneces: 1, indexes: ind}
			}
		}11
	}
	for k, v := range occuranceMap {

		if v.numOfOccureneces == 1 {
			return k, v.indexes[0], true
		}
	}
	res := []byte(".")[0]
	return res, [2]int{0, 0}, false

}
func containsInSudokuList(l []byte, val byte) bool {
	for _, el := range l {
		if el == val {
			return true
		}
	}
	return false
}

func containsInMatrix(board [][]byte, el byte) bool {
	for _, r := range board {
		if containsInSudokuList(r, el) {
			return true
		}
	}
	return false
}

func findElementThatDoesNotBelong(board [][]byte, index [2]int) (byte, [2]int, bool) {
	rowIntervals := make(map[[2]int][]byte)
	columnIntervals := make(map[[2]int][]byte)
	squareIntervals := make(map[[2]int][]byte)
	i, j := index[0], index[1]
	length := len(board)
	squareIndexes := getAllIndexesForElementsInSquare(index, int(math.Sqrt(float64(length))))

	for k := 0; k < length; k++ {
		if string(board[i][k]) != "." {
			continue
		}
		rowIntervals[[2]int{i, k}] = getPossibleValues(board, [2]int{i, k})
	}

	if el, ind, ok := determineInIntervals(rowIntervals); ok {
		fmt.Println("determined in interval row:", string(el), ind, ok)
		return el, ind, ok
	}

	for k := 0; k < length; k++ {
		if string(board[k][j]) != "." {
			continue
		}
		columnIntervals[[2]int{k, j}] = getPossibleValues(board, [2]int{k, j})
	}

	if el, ind, ok := determineInIntervals(columnIntervals); ok {
		fmt.Println("determined in interval column:", columnIntervals, string(el), ind, ok)
		return el, ind, ok
	}
	for _, i := range squareIndexes {
		if string(board[i[0]][i[1]]) != "." {
			continue
		}
		squareIntervals[i] = getPossibleValues(board, i)
	}

	if el, ind, ok := determineInIntervals(squareIntervals); ok {
		fmt.Println("determined in interval square:", string(el), ind, ok)
		return el, ind, ok
	}
	res_ind := [2]int{0, 0}
	return []byte(".")[0], res_ind, false

}

func getSquareForElement(board [][]byte, index [2]int) []byte {
	square := make([]byte, 0)
	N := int(math.Sqrt(float64(len(board))))
	upperBoundRow := N*(index[0]/N) + N
	lowerBoundRow := N * (index[0] / N)
	upperBoundColumn := N*(index[1]/N) + N
	lowerBoundColumn := N * (index[1] / N)
	i := lowerBoundRow

	for i < upperBoundRow {
		square = append(square, board[i][lowerBoundColumn:upperBoundColumn]...)
		i++
	}
	return square

}

func getColumn(board [][]byte, j int) []byte {
	column := []byte{}
	for i := 0; i < len(board); i++ {
		column = append(column, board[i][j])
	}
	return column
}

func getPossibleValues(board [][]byte, index [2]int) []byte {
	i, j := index[0], index[1]
	row := board[i]
	column := getColumn(board, j)
	square := getSquareForElement(board, index)
	possibleValues := []byte{}
	for _, pv := range sudokuValues {
		if containsInSudokuList(row, pv) || containsInSudokuList(column, pv) || containsInSudokuList(square, pv) {
			continue
		}
		possibleValues = append(possibleValues, pv)
	}
	return possibleValues

}
func getPossibleValuesAfterAssume(board [][]byte, assumed_decision map[[2]int]byte, index [2]int) []byte {
	boardCopy := board
	for ind, v := range assumed_decision {
		boardCopy[ind[0]][ind[1]] = v
	}
	return getPossibleValues(boardCopy, index)
}

func makeNewDecision(board [][]byte, decisions []map[[2]int]byte) {
	// b := buildMapFromBoard(board)
	// currentDecisions
	// for i, dec := range b 	// }
	if len(decisions) == 0 {
		initialBoard := BuildMapFromBoardForPossibleValues(board)

		for k, v := range initialBoard {
			for _, i := range v {
				decisions = append(decisions, map[[2]int]byte{k: i})
			}
		}
	}

}

func traverseAndSolveWithDecisionTree(board [][]byte) {

	var decisions []map[[2]int]byte

	initialBoard := BuildMapFromBoardForPossibleValues(board)

	for k, v := range initialBoard {
		for _, i := range v {
			decisions = append(decisions, map[[2]int]byte{k: i})
		}
	}

	boardCopy := board
	for _, d := range decisions {
		for ind, val := range d {
			boardCopy[ind[0]][ind[1]] = val
			for traverseAndSolveWithoutDecisionTree(boardCopy) {

			}
		}
	}

}

func traverseAndSolveWithoutDecisionTree(board [][]byte) bool {
	boardCopy := board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if string(board[i][j]) != "." {
				continue
			}
			if len(getPossibleValues(board, [2]int{i, j})) == 1 {
				board[i][j] = getPossibleValues(board, [2]int{i, j})[0]
				fmt.Println("evaluation", string(board[i][j]), []int{i, j})
			}

		}

	}
	if containsInMatrix(board, []byte(".")[0]) {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board); j++ {
				if string(board[i][j]) != "." {
					continue
				}
				pv, index, ok := findElementThatDoesNotBelong(board, [2]int{i, j})
				if ok {
					board[index[0]][index[1]] = pv
					fmt.Println("determination", string(board[i][j]), index, string(pv))
				}

			}

		}

	}
	return !compareTwoBoards(board, boardCopy)
}

func solveSudoku(board [][]byte) {
	for traverseAndSolveWithoutDecisionTree(board) {
		continue
	}

}
