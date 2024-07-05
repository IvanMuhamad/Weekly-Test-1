package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{1, 4, 8, 9}
	nums3 := []int{9, 9, 9, 9}

	output1 := plusOne(nums1)
	output2 := plusOne(nums2)
	output3 := plusOne(nums3)
	fmt.Println(output1)
	fmt.Println(output2)
	fmt.Println(output3)
}

func plusOne(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i] += 1
			return nums
		}
		nums[i] = 0
	}
	return append([]int{1}, nums...)
}
