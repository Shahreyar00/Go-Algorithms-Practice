package main

import "fmt"

func checkSubarray(nums []int, k int) bool {
	remainderMap := make(map[int]int)
	remainderMap[0] = -1
	sum := 0

	for i, num := range nums {
		sum += num
		remainder := sum % k
		if remainder < 0 {
			remainder += k
		}

		if prevIndex, exists := remainderMap[remainder]; exists {
			if i-prevIndex > 1 {
				return true
			}
		} else {
			remainderMap[remainder] = i
		}
	}

	return false
}

func main() {
	var n, k int
	fmt.Println("Enter the size of the array: ")
	fmt.Scan(&n)

	nums := make([]int, n)
	fmt.Println("Enter the elements of the array: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println("Enter the value of k: ")
	fmt.Scan(&k)

	result := checkSubarray(nums, k)
	fmt.Println("Output: ", result)
}
