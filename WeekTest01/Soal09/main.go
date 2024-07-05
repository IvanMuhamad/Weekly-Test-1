package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 4, 4, 3, 3, 2, 4}
	nums2 := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}

	frequency1 := countFrequentElement(nums1)
	for num, count := range frequency1 {
		fmt.Printf("%d = %d\n", num, count)
	}
	fmt.Println("")
	frequency2 := countFrequentElement(nums2)
	for num, count := range frequency2 {
		fmt.Printf("%d = %d\n", num, count)
	}
}

func countFrequentElement(nums []int) map[int]int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num] += 1
	}

	return count
}
