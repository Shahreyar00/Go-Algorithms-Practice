package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := low + (high-low)/2
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				low = mid + 2
			} else {
				high = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				low = mid + 1
			} else {
				high = mid
			}
		}
	}

	return nums[low]
}

func main() {
	var n int
	fmt.Println("Enter the number of elements in the array:")
	fmt.Scan(&n)

	nums := make([]int, n)
	fmt.Println("Enter the elements of the array (sorted):")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	result := singleNonDuplicate(nums)
	fmt.Println("The single non-duplicate element is: ", result)
}
