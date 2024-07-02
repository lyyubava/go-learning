package main

import "fmt"

func containsString(arr []string, el string) bool {
	for _, c := range arr {
		if c == el {
			return true
		}
	}
	return false
}

func isQueueEmpty(queue *[]string) bool {
	if len(*queue) == 0 {
		return true
	}
	return false
}

func enqueue(queue *[]string, el string) {
	*queue = append(*queue, el)
}

func dequeue(queue *[]string) string {
	if isQueueEmpty(queue) {
		return ""
	}
	if len(*queue) == 1 {
		el := (*queue)[0]
		*queue = make([]string, 0)
		return el

	}
	el := (*queue)[0]
	*queue = (*queue)[1:]
	return el

}

func lengthOfLongestSubstring(s string) []string {
	queue := make([]string, 0)
	tmpQueue := make([]string, 0)
	resultQueue := make([]string, 0)
	for _, char := range s {
		enqueue(&queue, string(char))
	}
	for !isQueueEmpty(&queue) {
		symbol := dequeue(&queue)
		for !containsString(tmpQueue, symbol) {
			enqueue(&tmpQueue, symbol)
			if !isQueueEmpty(&queue) {
				symbol = dequeue(&queue)
				continue
			}
			break

		}
		if len(tmpQueue) > len(resultQueue) {
			resultQueue = tmpQueue
		}
		enqueue(&tmpQueue, symbol)
		for dequeue(&tmpQueue) != symbol {
			continue
		}

	}
	return resultQueue

}

func main() {
	var s string
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(len(lengthOfLongestSubstring(s)))

}
