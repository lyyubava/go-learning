package main

import "fmt"

type stack []int
type matrix [][]int

func (s *stack) pop() int {
	if len(*s) == 0 {
		return 0
	}
	el := (*s)[len((*s))-1]
	(*s) = (*s)[0 : len((*s))-1]
	return el
}

func (s *stack) push(el int) {
	new_s := make([]int, 0)
	new_s = append(new_s, el)
	new_s = append(new_s, (*s)...)
	*s = *&new_s
}

func generateMatrix(n int) [][]int {
	s := stack{}
	for i := 1; i <= n*n; i++ {
		s.push(i)
	}
	m := matrix{}
	m.fill(-1, n)
	rowStart := 0
	rowEnd := n - 1
	columnStart := n - 1
	columnEnd := 0
	for columnEnd <= columnStart && rowEnd >= rowStart {
		for i := columnEnd; i <= columnStart; i++ {
			m[rowStart][i] = s.pop()
		}
		rowStart += 1
		for j := rowStart; j <= rowEnd; j++ {
			m[j][columnStart] = s.pop()
		}
		columnStart -= 1
		for k := columnStart; k >= columnEnd; k-- {
			m[rowEnd][k] = s.pop()
		}
		rowEnd -= 1
		for x := rowEnd; x >= rowStart; x-- {
			m[x][columnEnd] = s.pop()
		}

		columnEnd += 1
	}
	return m
}

func (m *matrix) fill(v int, n int) {

	for x := 0; x < n; x++ {
		row := make([]int, 0)
		for i := 0; i < n; i++ {
			row = append(row, v)

		}
		(*m) = append(*m, row)
	}

}
func main() {
	m := generateMatrix(2)
	fmt.Println(m)
}
