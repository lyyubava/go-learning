package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head
	for i := 0; i < getLen(head)-n-1; i++ {
		curr = curr.Next

	}
	if n == getLen(head) && getLen(head) > 1 {
		head = curr.Next
		return head
	}
	if n == getLen(head) && getLen(head) == 1 {
		head = nil
		return head
	}
	curr.Next = curr.Next.Next
	return head
}

func main() {

	// head = [1,2,3,4,5], n = 2

	// n5 := ListNode{Val: 5, Next: nil}
	// n4 := ListNode{Val: 4, Next: &n5}
	// n3 := ListNode{Val: 3, Next: &n4}
	// n2 := ListNode{Val: 2, Next: &n3}
	// n1 := ListNode{Val: 1, Next: &n2}

	j1 := ListNode{Val: 1, Next: nil}
	N := 1
	newHead := removeNthFromEnd(&j1, N)
	fmt.Println(getLen(newHead))

	//fmt.Println(removeNthFromEnd(&n1, N))
}
