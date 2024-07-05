package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	words := []string{"this", "is", "a", "kanoha"}
	except := []string{"is", "a"}

	result := capitalize(words, except)
	fmt.Println(result)
}

func checkException(word string, except []string) bool {
	for _, v := range except {
		if v == word {
			return true
		}
	}
	return false
}
func capitalize(words []string, except []string) []string {
	output := make([]string, len(words))

	for i, word := range words {
		if checkException(word, except) {
			output[i] = word
		} else {
			output[i] = cases.Title(language.Und).String(word)
		}
	}

	return output
}
