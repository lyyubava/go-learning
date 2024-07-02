package main

import "fmt"

func binarySearch(nums []int, low, high, target int) int {
	if high >= low {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			return binarySearch(nums, low, mid-1, target)
		} else {
			return binarySearch(nums, mid+1, high, target)
		}
	} else {
		return -1
	}

}

func traverse(nums []int, low, high, target int) int {

	if high >= low {
		mid := (low + high) / 2
		if nums[low] <= nums[mid] {
			if nums[mid] >= target && nums[low] <= target {
				return binarySearch(nums, low, mid+1, target)
			} else {
				return traverse(nums, mid+1, high, target)
			}
		} else if nums[mid] <= nums[high] {
			if nums[mid] <= target && nums[high] >= target {
				return binarySearch(nums, mid-1, high, target)
			} else {
				return traverse(nums, low, mid-1, target)
			}
		}

	}
	return -1
}
func search(nums []int, target int) int {
	return traverse(nums, 0, len(nums)-1, target)
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 2
	res := search(nums, target)
	fmt.Println(res)
}
