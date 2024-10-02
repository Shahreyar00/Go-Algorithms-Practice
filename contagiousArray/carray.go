package main

import (
	"fmt"
)

func findMaxLength(nums []int) int {
	sumMap := make(map[int]int)
	sumMap[0] = -1

	maxLength := 0
	sum := 0

	for i, num := range nums {
		if num == 0 {
			sum += -1
		} else {
			sum += 1
		}

		if prevIndex, exists := sumMap[sum]; exists {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			sumMap[sum] = i
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Println("Enter the number of elements in the binary array:")
	fmt.Scan(&n)

	nums := make([]int, n)
	fmt.Println("Enter the elements of the binary array (0s and 1s only):")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	result := findMaxLength(nums)
	fmt.Println("The maximum length of a contiguous subarray with equal number of 0s and 1s is:", result)
}
