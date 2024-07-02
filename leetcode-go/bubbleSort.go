package main

import (
	"fmt"
	"os"
	"strconv"
)

func Swap(sli []int, i int) {
	tmp := sli[i+1]
	sli[i+1] = sli[i]
	sli[i] = tmp
}

func IterateAndSwap(sli []int) int {
	var i, swaps int
	i, swaps = 0, 0

	for i < len(sli)-1 {
		if sli[i] > sli[i+1] {
			Swap(sli, i)
			swaps++
		}
		i++
	}
	return swaps
}

func BubbleSort(sli []int) {
	// here we just identify if array is already sorted or needs to go through another iteration of swapping
	for true {
		if IterateAndSwap(sli) == 0 {
			break
		} else {
			IterateAndSwap(sli)
		}
	}
}

func main() {
	var nums []int
	var num string
	var n int
	fmt.Print("Enter the number of element in sequence: ")
	fmt.Scan(&n)

	i := 0
	fmt.Print("Enter the sequence of up to 10 integers: ")
	for i < n {
		fmt.Scan(&num)
		if num == "" {
			break
		}
		num, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("The program perfoms sorting on integers!!")
			os.Exit(1)
		}
		nums = append(nums, num)
		i++
	}
	BubbleSort(nums)
	fmt.Println(nums)
}
