package main

import "fmt"

func findMinMoves(machines []int) int {
	n := len(machines)
	totalDresses := 0

	for _, dresses := range machines {
		totalDresses += dresses
	}

	if totalDresses%n != 0 {
		return -1
	}

	target := totalDresses / n
	maxMoves := 0
	runningBalance := 0

	for _, dresses := range machines {
		diff := dresses - target
		runningBalance += diff
		maxMoves = max(maxMoves, abs(runningBalance), diff)
	}

	return maxMoves
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Print("Enter the number of machines: ")
	fmt.Scan(&n)

	machines := make([]int, n)
	fmt.Printf("Enter the number of dresses in each machine: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&machines[i])
	}

	result := findMinMoves(machines)
	fmt.Printf("Minimum moves needed: %d\n", result)
}
