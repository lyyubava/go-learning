package main

import "fmt"

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
	value     int
	parent    *Node
	childrens *[]Node
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
	ch := []Node{}
	r := Node{
		value:     1,
		parent:    nil,
		childrens: &ch,
	}

	// insert(2, &r)
	ch3 := insert(3, &r)
	_ = insert(9, &ch3)
	ch10 := insert(10, &ch3)
	_ = insert(11, &ch10)
	_ = insert(2, &r)

	// fmt.Println((*((*(r.childrens))[0].childrens))[0].childrens)
	res := getAllChildrensWithoutChildrens(&r)
	fmt.Println(res)

}
func insert(val int, parent *Node) Node {
	child := Node{
		value:     val,
		parent:    parent,
		childrens: &([]Node{}),
	}
	*parent.childrens = append(*parent.childrens, child)
	return child
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
func getAllPossibleValues(board [][]byte, index [2]int) []byte {
	// return a map where key is index and value is all possible values for thet index
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

func getAllPossibleValues(board [][]byte, index){

}
func makeDecision(node Node) {
	if 
}