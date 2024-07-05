package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 4, 7, 8}
	nums2 := []int{2, 3, 7, 9}

	nums3 := []int{1, 2, 7, 4, 7, 8}
	nums4 := []int{7, 7, 3, 2, 9}

	nums5 := []int{2, 4, 6, 8}
	nums6 := []int{1, 3, 5, 7, 9}

	output1 := sameElement(nums1, nums2)
	output2 := sameElement(nums3, nums4)
	output3 := sameElement(nums5, nums6)
	fmt.Println(output1)
	fmt.Println(output2)
	fmt.Println(output3)
}

func sameElement(nums1 []int, nums2 []int) []int {
	check := make(map[int]bool)
	result := []int{}

	for _, num := range nums1 {
		check[num] = true
	}

	for _, num := range nums2 {
		if check[num] {
			result = append(result, num)
			delete(check, num)
		}
	}

	return result
}
