package main

import (
	"fmt"
	"math"
)

func findRotateSteps(ring string, key string) int {
	n := len(ring)
	m := len(key)

	pos := make(map[byte][]int)
	for i := 0; i < n; i++ {
		pos[ring[i]] = append(pos[ring[i]], i)
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	for _, j := range pos[key[0]] {
		dp[0][j] = min(j, n-j) + 1
	}

	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]] {
			for _, k := range pos[key[i-1]] {
				step := min(abs(j-k), n-abs(j-k))
				dp[i][j] = min(dp[i][j], dp[i-1][k]+step+1)
			}
		}
	}

	result := math.MaxInt32
	for _, j := range pos[key[m-1]] {
		result = min(result, dp[m-1][j])
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var ring, key string

	fmt.Print("Enter the ring string: ")
	fmt.Scan(&ring)

	fmt.Print("Enter the key string: ")
	fmt.Scan(&key)

	fmt.Println("Minimum number of steps:", findRotateSteps(ring, key))
}
