package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func fill(ln *ListNode, sli []int) *ListNode {
	head := &ListNode{Val: -1}
	curr := head
	for _, el := range sli {
		if curr.Val == -1 {

			curr.Val = el

		} else {
			curr.Next = &ListNode{Val: el}
			curr = curr.Next

		}

	}
	return head
}

func addNode(head *ListNode, el int) {
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = &ListNode{Val: el}

}
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
}

func getLength(head *ListNode) int {
	curr := head
	len := 0
	for curr != nil {
		len += 1
		curr = curr.Next
	}
	return len
}

func equalizeLists(l1 *ListNode, l2 *ListNode) {
	desiredLength := 0
	l := &ListNode{}

	if getLength(l1) > getLength(l2) {
		desiredLength = getLength(l1)
		l = l2
	} else {
		desiredLength = getLength(l2)
		l = l1
	}
	curr := getLength(l)
	for curr < desiredLength {
		addNode(l, 0)
		curr++
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := make([]int, 0)
	equalizeLists(l1, l2)
	offset := 0
	for l1 != nil {
		resNode := l1.Val + l2.Val + offset
		if resNode >= 10 {
			offset = 1
			resNode = resNode - 10
		} else {
			offset = 0
		}

		res = append(res, resNode)
		l1 = l1.Next
		l2 = l2.Next
	}
	if offset == 1 {
		res = append(res, 1)
	}
	return fill(&ListNode{}, res)
}
func main() {
	head1 := fill(&ListNode{}, []int{9, 9, 9, 9, 9, 9, 9})
	head2 := fill(&ListNode{}, []int{9, 9, 9, 9})
	h := addTwoNumbers(head1, head2)
	printList(h)

}
