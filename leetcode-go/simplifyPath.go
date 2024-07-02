package main

import (
	"fmt"
	"strings"
)

func pop(s *[]string) string {
	l := len(*s)
	if l == 0 {
		return ""
	}
	last := (*s)[l-1]
	*s = (*s)[:l-1]
	return last
}

func push(s *[]string, el string) string {
	*s = append(*s, el)
	return el
}

func simplifyPath(path string) string {

	str := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, s := range str {
		if s == "" || s == "." {
			continue
		} else if s == ".." {
			pop(&stack)
			continue
		}
		push(&stack, s)
	}
	if len(stack) == 0 {
		return "/"
	}

	return "/" + strings.Join(stack, "/")

}

func main() {
	var path string
	path = "/a//b////c/d//././/.."
	res := simplifyPath(path)
	fmt.Println(res)
}
