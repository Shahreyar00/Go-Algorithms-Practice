package main

import "fmt"

func findKthNumber(n, k int) int {
	curr := 1
	k--

	for k > 0 {
		steps := calcSteps(n, curr, curr+1)
		if steps <= k {
			curr++
			k -= steps
		} else {
			curr *= 10
			k--
		}
	}

	return curr
}

func calcSteps(n, curr, next int) int {
	steps := 0
	for curr <= n {
		steps += min(n+1, next) - curr
		curr *= 10
		next *= 10
	}

	return steps
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, k int
	fmt.Println("Enter the value of n and k:")
	_, err := fmt.Scanf("%d %d", &n, &k)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	result := findKthNumber(n, k)
	fmt.Printf("The %d-th smallest number in lexicographical order is: %d\n", k, result)
}
