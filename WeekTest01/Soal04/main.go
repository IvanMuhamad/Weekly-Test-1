package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 4, 4, 4}
	output := removeDuplicate(nums)
	fmt.Println(output)
}

func removeDuplicate(nums []int) []int {
	check := make(map[int]bool)
	output := []int{}

	for _, num := range nums {
		if !check[num] {
			check[num] = true
			output = append(output, num)
		}
	}
	return output
}
