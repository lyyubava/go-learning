package main

import "fmt"

type Node struct {
	value     string
	parent    *Node
	childrens *[]Node
}
type LinkedListNode struct {
	data []string
	next *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
}

var parentTree Node = Node{value: "1", parent: nil, childrens: &[]Node{}}
var linkedList LinkedList = LinkedList{head: nil}

// func buildTree(combinations LinkedList) {
// 	parents := combinations.head
// 	for parents != nil {
// 		childrens := parents.next
// 		for _, parent := range parents.data {
// 			for _, child := range childrens.data {

// 				insert(child,

// 				]]]]])
// 			}

// 		}

// 	}
// }

// func createBaseTree(parents )

func createCombinations() {

}
func main() {

	insert("3", &parentTree)
	fmt.Println(parentTree.childrens)

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

	return []string{}
}

func insert(val string, parent *Node) Node {
	child := Node{
		value:     val,
		parent:    parent,
		childrens: &([]Node{}),
	}
	fmt.Println(parent.childrens)
	*parent.childrens = append(*parent.childrens, child)
	return child
}
