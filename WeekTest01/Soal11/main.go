package main

import "fmt"

func main() {
	fruits1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	fruits2 := []string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}

	fmt.Println(fruitOperation(fruits1, fruits2, 1))
	fmt.Println(fruitOperation(fruits1, fruits2, 2))
}

func fruitOperation(fruits1 []string, fruits2 []string, operationType int) []string {
	switch operationType {
	case 1:
		return same(fruits1, fruits2)
	case 2:
		return difference(fruits1, fruits2)
	default:
		return []string{}
	}
}

func same(fruits1 []string, fruits2 []string) []string {
	check := make(map[string]bool)
	output := []string{}

	for _, fruit := range fruits1 {
		check[fruit] = true
	}

	for _, fruit := range fruits2 {
		if check[fruit] {
			output = append(output, fruit)
			delete(check, fruit)
		}
	}

	return output
}

func difference(fruits1 []string, fruits2 []string) []string {
	check1 := make(map[string]bool)
	check2 := make(map[string]bool)
	output := []string{}

	for _, fruit := range fruits1 {
		check1[fruit] = true
	}

	for _, fruit := range fruits2 {
		check2[fruit] = true
	}

	for _, fruit := range fruits1 {
		if !check2[fruit] {
			output = append(output, fruit)
		}
	}

	for _, fruit := range fruits2 {
		if !check1[fruit] {
			output = append(output, fruit)
		}
	}

	return output
}
