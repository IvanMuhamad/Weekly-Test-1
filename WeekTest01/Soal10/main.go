package main

import (
	"fmt"
)

func addDigits(nums1 []int, nums2 []int) []int {
	num1 := len(nums1)
	num2 := len(nums2)
	maxLen := max(num1, num2) //pilih array terpanjang untuk array hasil
	result := make([]int, maxLen+1)

	store := 0
	i := 0
	for i < num1 || i < num2 {
		sum := store
		if i < num1 {
			sum += nums1[num1-1-i]
		}
		if i < num2 {
			sum += nums2[num2-1-i]
		}
		store = sum / 10
		result[maxLen-i] = sum % 10
		i += 1
	}

	if store > 0 {
		result[0] = store
		return result
	}

	return result[1:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6}

	nums3 := []int{9, 2, 7}
	nums4 := []int{1, 3, 5}

	result := addDigits(nums1, nums2)
	fmt.Println(result)

	result2 := addDigits(nums3, nums4)
	fmt.Println(result2)
}
