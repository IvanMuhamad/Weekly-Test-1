package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	word1 := "Otto"
	word2 := "Toto"

	fmt.Println(isAnagram(word1, word2))
}

func sortString(words string) string {
	char := []rune(words)

	sort.Slice(char, func(i, j int) bool {
		return char[i] < char[j]
	})

	return string(char)
}

func isAnagram(word1 string, word2 string) bool {
	word1 = strings.ReplaceAll(strings.ToLower(word1), " ", " ")
	word2 = strings.ReplaceAll(strings.ToLower(word2), " ", " ")

	if sortString(word1) == sortString(word2) {
		return true
	} else {
		return false
	}
}
