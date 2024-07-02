package main

import (
	"fmt"
	"strings"
)

func main() {

	value := make([][]string, 0)
	value = append(value, []string{"a", "b", "c"})
	value = append(value, []string{"d", "e", "f"})
	value = append(value, []string{"g", "h", "i"})
	combinations := createCombinations(value)
	fmt.Println(len(combinations))
	res := letterCombinations("2")
	fmt.Println(res)

}

func mergeTwoCombinations(combination1, combination2 [][]string) [][]string {
	res := make([][]string, 0)
	if len(combination1) == 0 {
		res = combination2
		return res
	}
	if len(combination2) == 0 {
		res = combination1
		return res
	}
	for _, c1 := range combination1 {
		for _, c2 := range combination2 {
			comb := make([]string, 0)
			comb = append(comb, c1...)
			comb = append(comb, c2...)
			res = append(res, comb)
		}
	}
	return res
}
func createCombinations(combinations [][]string) [][]string {
	res := make([][]string, 0)
	if len(combinations) == 1 {
		for _, combination := range combinations {
			for _, val := range combination {
				res = append(res, []string{val})
			}

		}
		return res
	}

	for _, combination := range combinations {
		if len(res) == 0 {
			res = append(res, combination)
			res = createCombinations(res)
			continue
		}
		tmp := make([][]string, 0)
		tmp = append(tmp, combination)
		combs := createCombinations(tmp)
		res = mergeTwoCombinations(res, combs)
	}

	return res

}
func letterCombinations(digits string) []string {
	m := make(map[string][]string, 0)
	m["2"] = []string{"a", "b", "c"}
	m["3"] = []string{"d", "e", "f"}
	m["4"] = []string{"g", "h", "i"}
	m["5"] = []string{"j", "k", "l"}
	m["6"] = []string{"m", "n", "o"}
	m["7"] = []string{"p", "q", "r", "s"}
	m["8"] = []string{"t", "u", "v"}
	m["9"] = []string{"w", "x", "y", "z"}
	res := make([]string, 0)
	combinations := make([][]string, 0)
	for _, digit := range strings.Split(digits, "") {
		combinations = append(combinations, m[digit])
	}
	combinationsRes := createCombinations(combinations)
	for _, combination := range combinationsRes {
		res = append(res, strings.Join(combination, ""))
	}
	return res
}
