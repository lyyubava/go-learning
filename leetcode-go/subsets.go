package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

func getLen(head *ListNode) int {
	len := 0
	tmp := head
	for tmp != nil {
		len += 1
		tmp = tmp.Next
	}
	return len
}

func makeListFromSlice(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := ListNode{Val: arr[0], Next: nil, Prev: nil}

	tmp := &head

	for i := 1; i < len(arr); i++ {
		curr := ListNode{Val: arr[i], Next: nil, Prev: nil}
		tmp.Next = &curr
		tmp = &curr
	}
	prev := &head
	tmp = head.Next

	for tmp != nil {
		tmp.Prev = prev
		tmp = tmp.Next
		prev = prev.Next
	}

	return &head

}

func covertListOfListNOdesToSlice(list []*ListNode) []int {
	res := make([]int, 0)
	for _, el := range list {
		res = append(res, el.Val)
	}
	return res
}

func checkIfIthElementAfterElementIsNil(el *ListNode, i int) bool {
	tmp := el
	if tmp == nil {
		return true
	}
	for j := 0; j < i; j++ {

		tmp = tmp.Next
		if tmp == nil {
			return true
		}

	}
	return false
}

func fillListFromNthElementWithNexts(initPart []*ListNode, curr *ListNode, currIndex int) [][]int {
	result := make([][]int, 0)
	tmp := curr.Next
	res := make([]int, 0)
	front := initPart[:currIndex]
	res = append(res, covertListOfListNOdesToSlice(front)...)
	res = append(res, curr.Val)
	for j := 0; j < len(initPart)-currIndex-1; j++ {
		res = append(res, tmp.Val)
		tmp = tmp.Next
	}
	return res

}

func subsets(nums []int, N int) []int {
	var res []int
	initLinkedListhead := makeListFromSlice(nums)
	init := make([]*ListNode, 0)
	tmp := initLinkedListhead
	for i := 0; i < N; i++ {
		init = append(init, tmp)
		tmp = tmp.Next
	}
	if !checkIfIthElementAfterElementIsNil(init[0], 1) {
		return fillListFromNthElementWithNexts(init, init[0], 0)
	}
	res = []int{0}
	return res

}

func main() {
	var x []int = []int{1, 2, 3}
	r := subsets(x, 2)

	fmt.Println(r)

}
