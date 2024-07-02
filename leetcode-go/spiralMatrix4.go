package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type matrix [][]int

func (matrix *matrix) fill(v int, m int, n int) {

	for x := 0; x < m; x++ {
		row := make([]int, 0)
		for i := 0; i < n; i++ {
			row = append(row, v)

		}
		(*matrix) = append(*matrix, row)
	}

}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	matrix := matrix{}
	matrix.fill(-1, m, n)
	rowStart := 0
	rowEnd := m - 1
	columnStart := n - 1
	columnEnd := 0
	curr := head
	for columnEnd <= columnStart && rowEnd >= rowStart {
		for i := columnEnd; i <= columnStart; i++ {
			if curr != nil {
				matrix[rowStart][i] = curr.Val
			} else {
				break
			}
			curr = curr.Next

		}
		rowStart += 1
		for j := rowStart; j <= rowEnd; j++ {

			if curr != nil {
				matrix[j][columnStart] = curr.Val
			} else {
				break
			}
			curr = curr.Next
		}
		columnStart -= 1
		for k := columnStart; k >= columnEnd; k-- {
			if curr != nil {
				matrix[rowEnd][k] = curr.Val
			} else {
				break
			}
			curr = curr.Next
		}
		rowEnd -= 1
		for x := rowEnd; x >= rowStart; x-- {
			if curr != nil {
				matrix[x][columnEnd] = curr.Val
			} else {
				break
			}
			curr = curr.Next
		}

		columnEnd += 1
	}

	return matrix
}
