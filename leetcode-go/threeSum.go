package main

import (
	"fmt"
	"sort"
)

func contains(nums []int, el int) bool {
	for _, val := range nums {
		if int(val) == el {
			return true
		}
	}
	return false
}

func containsFindIndex(nums []int, el int) int {
	for pos, val := range nums {
		if int(val) == el {
			return pos
		}
	}
	return -1
}

func compareTwoArrays(num1 []int, num2 []int) bool {
	if len(num1) != len(num2) {
		return false
	}
	tmp1 := make([]int, 0)
	tmp2 := make([]int, 0)
	for i := 0; i < len(num1); i++ {

		for j := 0; j < len(num2); j++ {
			if !contains(num2, num1[i]) {
				return false
			}
			ind1 := containsFindIndex(num1, num1[i])
			ind2 := containsFindIndex(num2, num1[i])

			tmp1 = append(tmp1, num1[:ind1]...)
			tmp1 = append(tmp1, num1[ind1+1:]...)

			tmp2 = append(tmp2, num2[:ind2]...)
			tmp2 = append(tmp2, num2[ind2+1:]...)

			num1 = tmp1
			num2 = tmp2

			tmp1 = make([]int, 0)
			tmp2 = make([]int, 0)

		}

	}
	return true

}

func containsBoth(nums []int, el int, el1 int) bool {
	if !contains(nums, el) || !contains(nums, el1) {
		return false
	}
	tmp := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == el {
			tmp = append(tmp, nums[:i]...)
			tmp = append(tmp, nums[i+1:]...)
			break
		}
	}
	if contains(tmp, el1) {
		return true
	}
	return false
}

func removeDuplicates(arrays [][]int) [][]int {
	uniqueArrays := make([][]int, 0)

	for _, arr := range arrays {
		isDuplicate := false

		// Check if the current array is a duplicate of any previously processed array
		for _, uniqueArr := range uniqueArrays {
			if compareTwoArrays(arr, uniqueArr) {
				isDuplicate = true
				break
			}
		}

		// If it's not a duplicate, add it to the unique arrays
		if !isDuplicate {
			uniqueArrays = append(uniqueArrays, arr)
		}
	}

	return uniqueArrays
}

func threeSum(nums []int) [][]int {
	arr := sort.IntSlice(nums)

	lessTnenZero := make([]int, 0)
	biggerTnenZero := make([]int, 0)
	res := make([][]int, 0)

	for i := 0; i < len(arr); i++ {
		if arr[i] >= 0 {
			biggerTnenZero = append(biggerTnenZero, arr[i])
			continue
		}
		lessTnenZero = append(lessTnenZero, arr[i])

	}
	if len(lessTnenZero) > 0 && len(biggerTnenZero) > 0 {
		for _, el := range lessTnenZero {
			triplet := make([]int, 0)
			for i := 0; i < len(biggerTnenZero); i++ {
				tmp := el + biggerTnenZero[i]
				if tmp > 0 {
					if containsBoth(lessTnenZero, 0-tmp, el) {
						triplet = append(triplet, el, biggerTnenZero[i], 0-tmp)
						res = append(res, triplet)
						triplet = make([]int, 0)
						continue
					}
				}
				if containsBoth(biggerTnenZero, 0-tmp, biggerTnenZero[i]) {
					triplet = append(triplet, el, biggerTnenZero[i], 0-tmp)
					res = append(res, triplet)
					triplet = make([]int, 0)
					continue
				}

			}
		}

		zeroAmount := 0
		for _, el := range biggerTnenZero {
			zeroAmount += el
		}
		if zeroAmount >= 3 {
			zeros := []int{0, 0, 0}
			res = append(res, zeros)
		}

	} else {
		arrayToTraverse := make([]int, 0)
		if len(lessTnenZero) > 0 {
			arrayToTraverse = lessTnenZero
		} else {
			arrayToTraverse = biggerTnenZero
		}
		zeroAmount := 0
		for _, el := range arrayToTraverse {
			if el == 0 {
				zeroAmount += 1
			}
		}
		if zeroAmount >= 3 {
			threeZeros := []int{0, 0, 0}

			res = append(res, threeZeros)
		}

	}
	return removeDuplicates(res)
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}
