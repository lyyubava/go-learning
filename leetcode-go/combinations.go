package main

import "fmt"

func main() {

	value := []string{"a", "b", "c", "d"}
	combs := createCombinations(value, 2)
	fmt.Println(combs)

}

func createCombinations(values []string, k int) [][]string {
	if k == 1 {
		res := make([][]string, 0)
		for _, value := range values {
			res = append(res, []string{value})
		}
		return res
	}
	res := make([][]string, 0)
	for i, value := range values {
		tail := make([]string, 0)
		tail = append(tail, values[:i]...)
		tail = append(tail, values[i+1:]...)
		combinations := createCombinations(tail, k-1)
		for _, combination := range combinations {
			combination = append(combination, value)
			res = append(res, combination)
		}
	}
	return res

}
