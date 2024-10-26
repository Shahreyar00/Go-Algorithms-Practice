package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func isSubsequence(s1, s2 string) bool {
	j := 0
	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s1[j] == s2[i] {
			j++
		}
	}
	return j == len(s1)
}

func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	for i, str := range strs {
		uncommon := true
		for j, other := range strs {
			if i != j && isSubsequence(str, other) {
				uncommon = false
				break
			}
		}

		if uncommon {
			return len(str)
		}
	}

	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the strings separated by spaces:")
	input, _ := reader.ReadString('\n')
	strs := strings.Fields(input)

	result := findLUSlength(strs)
	fmt.Println("The length of the longest uncommon subsequence is:", result)
}
