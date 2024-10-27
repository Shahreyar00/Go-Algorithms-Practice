package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the array elements separated by space:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strNums := strings.Split(input, " ")

	nums := make([]int, len(strNums))
	for i, str := range strNums {
		nums[i], _ = strconv.Atoi(str)
	}

	result := reversePairs(nums)
	fmt.Printf("Number of reverse pairs: %d\n", result)
}

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := (left + right) / 2
	count := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)
	count += countPairs(nums, left, mid, right)
	merge(nums, left, mid, right)
	return count
}

func countPairs(nums []int, left, mid, right int) int {
	count, j := 0, mid+1
	for i := left; i <= mid; i++ {
		for j <= right && nums[i] > 2*nums[j] {
			j++
		}
		count += j - (mid + 1)
	}
	return count
}

func merge(nums []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		temp[k] = nums[j]
		j++
		k++
	}
	for i := 0; i < len(temp); i++ {
		nums[left+i] = temp[i]
	}
}
