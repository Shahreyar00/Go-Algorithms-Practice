// DP implementation of Matrix Chain Multiplication problem
package main

import (
	"fmt"
	"math"
)

func matrixChainOrder(p []int) int {
	n := len(p) - 1
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			dp[i][j] = math.MaxInt32
			for k := i; k < j; k++ {
				q := dp[i][k] + dp[k+1][j] + p[i]*p[k+1]*p[j+1]
				if q < dp[i][j] {
					dp[i][j] = q
				}
			}
		}
	}

	return dp[0][n-1]
}

func main() {
	var n int
	fmt.Print("Enter the number of matrices: ")
	fmt.Scan(&n)

	p := make([]int, n+1)
	fmt.Println("Enter the dimensions of the matrices: ")
	for i := 0; i <= n; i++ {
		fmt.Printf("p[%d]: ", i)
		fmt.Scan(&p[i])
	}

	fmt.Printf("Minimum number of multiplications is %d\n", matrixChainOrder(p))

}
