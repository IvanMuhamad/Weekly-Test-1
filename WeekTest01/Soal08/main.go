package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{2, 4, 6}
	nums2 := []int{1, 3, 5, 7}

	fmt.Println(sliceOperation(nums1, nums2, 1))
	fmt.Println(sliceOperation(nums1, nums2, 2))
	fmt.Println(sliceOperation(nums1, nums2, 3))
	fmt.Println(sliceOperation(nums1, nums2, 4))
}

func sliceOperation(nums1 []int, nums2 []int, typeOperation int) []int {
	switch typeOperation {
	case 1:
		return difference1(nums1, nums2)
	case 2:
		return difference2(nums1, nums2)
	case 3:
		return union(nums1, nums2)
	case 4:
		return intersection(nums1, nums2)
	default:
		return []int{}
	}
}

func difference1(nums1 []int, nums2 []int) []int {
	check := make(map[int]bool)
	result := []int{}

	for _, num := range nums2 {
		check[num] = true
	}

	for _, num := range nums1 {
		if !check[num] {
			result = append(result, num)
		}
	}

	return result
}

func difference2(nums1 []int, nums2 []int) []int {
	check := make(map[int]bool)
	result := []int{}

	for _, num := range nums1 {
		check[num] = true
	}

	for _, num := range nums2 {
		if !check[num] {
			result = append(result, num)
		}
	}

	return result
}

func union(nums1 []int, nums2 []int) []int {
	check := make(map[int]bool)
	result := []int{}

	temp := append(nums1, nums2...)

	for _, num := range temp {
		if !check[num] {
			check[num] = true
			result = append(result, num)
		}
	}

	sort.Ints(result)
	return result
}

func intersection(nums1 []int, nums2 []int) []int {
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
