// DP implementation of LPS problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Approach 1
// func longestPalindromicSubsequence(s string) int {
// 	n := len(s)
// 	dp := make([][]int, n)
// 	for i := range dp {
// 		dp[i] = make([]int, n)
// 	}

// 	for i := 0; i < n; i++ {
// 		dp[i][i] = 1
// 	}

// 	for cl := 2; cl <= n; cl++ {
// 		for i := 0; i < n-cl+1; i++ {
// 			j := i + cl - 1
// 			if s[i] == s[j] && cl == 2 {
// 				dp[i][j] = 2
// 			} else if s[i] == s[j] {
// 				dp[i][j] = dp[i+1][j-1] + 2
// 			} else {
// 				dp[i][j] = utils.Max(dp[i][j-1], dp[i+1][j])
// 			}
// 		}
// 	}

// 	return dp[0][n-1]
// }

// Approach 2
func longestPalindromicSubsequence(s string) int {
	r := reverse(s)

	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == r[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[n][n]
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := longestPalindromicSubsequence(input)
	fmt.Printf("Length of the longest palindromic subsequence is: %d\n", result)
}
