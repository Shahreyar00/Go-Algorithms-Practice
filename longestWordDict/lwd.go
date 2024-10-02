package main

import (
	"fmt"
	"sort"
)

func isSubsequence(s, word string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(word) {
		if s[i] == word[j] {
			j++
		}
		i++
	}

	return j == len(word)
}

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) == len(dictionary[j]) {
			return dictionary[i] < dictionary[j]
		}
		return len(dictionary[i]) > len(dictionary[j])
	})

	for _, word := range dictionary {
		if isSubsequence(s, word) {
			return word
		}
	}

	return ""
}

func main() {
	var s string
	var n int
	fmt.Println("Enter the string:")
	fmt.Scan(&s)
	fmt.Println("Enter the number of words in the dictionary:")
	fmt.Scan(&n)

	dictionary := make([]string, n)
	fmt.Println("Enter the words in the dictionary:")
	for i := 0; i < n; i++ {
		fmt.Scan(&dictionary[i])
	}

	result := findLongestWord(s, dictionary)
	fmt.Println("The longest word is:", result)
}
