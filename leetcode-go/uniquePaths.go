package main

import "fmt"

type graph map[[2]int][][2]int

type q [][][2]int

func (q *q) isEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}
func (q *q) push(element [][2]int) {
	*q = append([][][2]int{element}, *q...)
}

func (q *q) pop() [][2]int {
	if q.isEmpty() {
		return [][2]int{}
	}
	popped := (*q)[0]
	(*q) = (*q)[1:]
	return popped
}

func (g graph) build(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			moved, _ := move(m, n, [2]int{i, j})
			for _, pos := range moved {
				g[[2]int{i, j}] = append(g[[2]int{i, j}], pos)
			}
		}
	}
}

func (g graph) findNumberOfPathsFromTwoGraphNodes(start [2]int, finish [2]int) int {
	num := 0
	q := q{}
	path := make([][2]int, 0)
	path = append(path, start)
	q.push(path)
	//fmt.Println(q)
	for !q.isEmpty() {
		//fmt.Println(q)
		p := q.pop()
		last := p[len(p)-1]
		if last == finish {
			num += 1
		}
		for i := 0; i < len(g[last]); i++ {
			newPath := path[:]
			newPath = append(newPath, g[last][i])
			q.push(newPath)
		}
	}
	return num
}

func move(m int, n int, pos [2]int) ([][2]int, error) {
	movedPos := make([][2]int, 0)
	moveDown, okDown := goDown(m, pos)
	moveRight, okRight := goRight(n, pos)
	err := fmt.Errorf("can't move anywhere left or right")
	if okDown == nil {
		movedPos = append(movedPos, moveDown)
		err = nil
	}
	if okRight == nil {
		movedPos = append(movedPos, moveRight)
		err = nil
	}
	return movedPos, err
}
func goDown(m int, pos [2]int) ([2]int, error) {
	var newPosition [2]int
	newPosition = pos
	err := fmt.Errorf("can't move anywhere down")
	if pos[0] < m-1 {
		newPosition[0] = pos[0] + 1
		err = nil
	} else {
		newPosition = [2]int{}
	}
	return newPosition, err

}

func goRight(n int, pos [2]int) ([2]int, error) {
	var newPosition [2]int
	newPosition = pos
	err := fmt.Errorf("can't move anywhere right")
	if pos[1] < n-1 {
		newPosition[1] = pos[1] + 1
		err = nil
	} else {
		newPosition = [2]int{}
	}
	return newPosition, err
}

func uniquePaths(m int, n int) int {
	g := graph{}
	g.build(m, n)
	fmt.Println("len", len(g))
	start := [2]int{0, 0}
	finish := [2]int{m - 1, n - 1}
	return g.findNumberOfPathsFromTwoGraphNodes(start, finish)

}
func main() {
	m, n := 23, 12
	fmt.Println(uniquePaths(m, n))
}
