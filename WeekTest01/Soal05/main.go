package main

import (
	"fmt"
)

func oddBeforeEvent(nums []int) []int {
	var genap []int
	var ganjil []int

	for _, num := range nums {
		if num%2 == 0 {
			genap = append(genap, num)
		} else {
			ganjil = append(ganjil, num)
		}
	}

	result := append(ganjil, genap...)
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(oddBeforeEvent(nums))
}
