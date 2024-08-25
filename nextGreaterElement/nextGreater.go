package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	stack := []int{}

	for i := range result {
		result[i] = -1
	}

	for i := 0; i < 2*n; i++ {
		num := nums[i%n]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = num
		}

		if i < n {
			stack = append(stack, i)
		}
	}

	return result
}

func main() {
	var n int
	fmt.Println("Enter the number of elements in the circular array:")
	fmt.Scan(&n)

	nums := make([]int, n)
	fmt.Println("Enter the elements of the circular array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	result := nextGreaterElements(nums)
	fmt.Println("Next greater elements:", result)
}
