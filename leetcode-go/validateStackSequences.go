package main

import "fmt"

func pop(s *[]int) int {
	l := len(*s)
	if l == 0 {
		return -1
	}
	last := (*s)[l-1]
	*s = (*s)[:l-1]
	return last
}

func push(s *[]int, el int) int {
	*s = append(*s, el)
	return el
}

func el_in_stack(s []int, el int) bool {
	result := false
	for _, val := range s {
		if val == el {
			result = true
			break
		}
	}
	return result
}
func validateStackSequences(pushed []int, popped []int) bool {
	var valid bool
	valid = true
	i := 0
	s := make([]int, 0)
	for _, pop_val := range popped {
		for !el_in_stack(s, pop_val) {
			push(&s, pushed[i])
			i += 1
		}
		if pop(&s) == pop_val {
			continue
		} else {
			valid = false
			return valid
		}

	}
	return valid
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	valid := validateStackSequences(pushed, popped)
	fmt.Println(valid)
}
