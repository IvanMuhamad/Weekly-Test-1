package main

import "fmt"

func main() {
	words1 := "bananas"
	words2 := "lalalamama"

	fmt.Println(removeDuplicate(words1))
	fmt.Println(removeDuplicate(words2))
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
