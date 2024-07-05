package main

import "fmt"

func main() {
	words1 := "bananas"
	words2 := "lalalamama"

	fmt.Println("Case 1: ", removeDuplicate(words1))
	fmt.Println("Case 2: ", removeDuplicate(words2))
}

func removeDuplicate(words string) string {
	check := make(map[rune]bool)
	output := []rune{}

	for _, char := range words {
		if !check[char] {
			check[char] = true
			output = append(output, char)
		}
	}
	return string(output)
}
