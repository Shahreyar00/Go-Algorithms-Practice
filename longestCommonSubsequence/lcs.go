// DP implementation of LCS problem
package main

import (
	"GoLang/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func longestCommonSubsequence(X, Y string) int {
	m := len(X)
	n := len(Y)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if X[i-1] == Y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the first string:")

	X, _ := reader.ReadString('\n')
	X = strings.TrimSpace(X)

	fmt.Println("Enter the second string:")

	Y, _ := reader.ReadString('\n')
	Y = strings.TrimSpace(Y)

	result := longestCommonSubsequence(X, Y)
	fmt.Printf("Length of the Longest Common Subsequence is: %d\n", result)
}
