package main

import (
	"math"
	"sort"
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

type Node struct {
	//value map[[2]int]int
	value     map[[2]int]byte
	parent    *Node
	childrens *[]Node
	solvable  bool
}

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
	//solveSudoku(arr)
	// ch := []Node{}
	// r := Node{
	// 	value:     1,
	// 	parent:    nil,
	// 	childrens: &ch,
	// }

	// // insert(2, &r)
	// ch3 := insert(3, &r)
	// _ = insert(9, &ch3)
	// ch10 := insert(10, &ch3)
	// _ = insert(11, &ch10)
	// _ = insert(2, &r)

	// fmt.Println((*((*(r.childrens))[0].childrens))[0].childrens)

	// res := getPossibleValuesByLength(arr)
	// fmt.Println(res)
}
func insert(val map[[2]int]byte, parent *Node) Node {
	child := Node{
		value:     val,
		parent:    parent,
		childrens: &([]Node{}),
	}
	*parent.childrens = append(*parent.childrens, child)
	return child
}

func deleteChildrensAndChildlessParents(child *Node) {
	curr := child
	parent := child.parent
	*(parent.childrens).del
}
func insertPossibleSolution(child *Node, board [][]byte) {
	boardCopy := board
	curr := child
	for true {
		for k := range curr.value {
			boardCopy[k[0]][k[1]] = child.value[k]
		}
		if curr.parent == nil {
			break
		}
	}
	pv, ok := getPossibleValuesByLength(boardCopy)
	if ok {
		keys := make([]int, 0)
		for k := range pv {
			keys = append(keys, k)
		}
		leastPv := sort.IntSlice(keys)[0]
		for i := range pv[leastPv] {
			valueToInsert := make(map[[2]int]byte, 0)
			valueToInsert[i] = pv[leastPv][i][0]
			insert(valueToInsert, child)
			break

		}
	} else {

	}

}
func getAllChildrensWithoutChildrens(root *Node) *[]Node {
	ch := []Node{}
	childrens := &ch
	res := childrens
	curr := root
	if len((*curr.childrens)) == 0 {

		*res = append(*res, *curr)
		return res
	} else {
		for _, ch := range *curr.childrens {
			r := getAllChildrensWithoutChildrens(&ch)
			*res = append(*res, *r...)

		}
	}
	return res
}

func containsInSudokuList(l []byte, val byte) bool {
	for _, el := range l {
		if el == val {
			return true
		}
	}
	return false
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

func getPossibleValuesByLength(board [][]byte) (map[int]map[[2]int][]byte, bool) {
	res := make(map[int]map[[2]int][]byte, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if string(board[i][j]) == "*" {
				continue
			}
			pv := getPossibleValues(board, [2]int{i, j})
			l := len(pv)
			if l == 0 {
				r := make(map[int]map[[2]int][]byte, 0)
				return r, false
			}
			_, ok := res[l]
			if !ok {
				tmp := make(map[[2]int][]byte, 0)
				tmp[[2]int{i, j}] = pv
				res[l] = tmp
			}
			res[l][[2]int{i, j}] = pv
		}
	}
	return res, true
}

func ifSolvable(node *Node, board [][]byte) bool {
	boardCopy := 
	var value map[[2]int]byte
	for node != nil {
		value = node.value
		for k, _ := range value {
			ind1, ind2 := k[0], k[1]
			board[ind1][ind2] = value[k]
		}

	}
	_, res := getPossibleValuesByLength(board)
	return res
}
func solve(board [][]byte) {
	var solution *Node

}

// func makeDecision(node Node) {
// 	if
// }
