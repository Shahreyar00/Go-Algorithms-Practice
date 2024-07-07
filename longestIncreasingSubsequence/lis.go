// DP implementation of LIS problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lenghtOfLis(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLenght := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		maxLenght = max(maxLenght, dp[i])
	}

	return maxLenght
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a list of integers separated by spaces:")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strNums := strings.Split(input, " ")

	var nums []int
	for _, strNum := range strNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("Invalid input. Please enter integers")
			return
		}
		nums = append(nums, num)
	}

	fmt.Println("Input array:", nums)
	result := lenghtOfLis(nums)
	fmt.Println("Length of the longest increasing subsequence is:", result)
}
