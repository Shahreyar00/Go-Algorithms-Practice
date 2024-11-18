package main

import "fmt"

func countArrangement(n int) int {
	var backtrack func(position int, used int) int
	backtrack = func(position int, used int) int {
		if position > n {
			return 1
		}

		count := 0
		for num := 1; num <= n; num++ {
			if used&(1<<num) != 0 {
				continue
			}

			if num%position == 0 || position%num == 0 {
				count += backtrack(position+1, used|(1<<num))
			}
		}
		return count
	}
	return backtrack(1, 0)
}

func main() {
	var n int
	fmt.Print("Enter the value of n: ")
	fmt.Scan(&n)

	result := countArrangement(n)
	fmt.Printf("The number of beautiful arrangements for n = %d is: %d\n", n, result)
}
