package main

import "fmt"

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	numFreq := make(map[int]int)
	for _, num := range nums {
		numFreq[num]++
	}

	count := 0
	for num := range numFreq {
		if k == 0 {
			if numFreq[num] > 1 {
				count++
			}
		} else {
			if _, exists := numFreq[num+k]; exists {
				count++
			}
		}
	}

	return count
}

func main() {
	var n, k int
	fmt.Println("Enter the number of elements in the array:")
	fmt.Scan(&n)

	nums := make([]int, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println("Enter the value of k:")
	fmt.Scan(&k)

	result := findPairs(nums, k)
	fmt.Printf("The number of unique k-diff pairs is: %d\n", result)
}
