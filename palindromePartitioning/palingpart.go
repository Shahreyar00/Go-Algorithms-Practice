package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(str string, start, end int) bool {
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func minCutsPalindromePartition(str string) int {
	n := len(str)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = i
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			if isPalindrome(str, j, i) {
				if j == 0 {
					dp[i] = 0
				} else {
					dp[i] = min(dp[i], dp[j-1]+1)
				}
			}
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a string:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := minCutsPalindromePartition(input)
	fmt.Println("Minimum cuts needed for palindrome partitioning:", result)
}
