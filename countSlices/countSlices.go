package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	var totalCount int

	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			diff := nums[j] - nums[i]
			countI := dp[i][diff]
			countJ := dp[i][diff]

			dp[j][diff] = countJ + countI + 1

			totalCount += countI
		}
	}

	return totalCount
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter numbers separated by commas: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numStrs := strings.Split(input, ",")
	nums := make([]int, len(numStrs))

	for i, s := range numStrs {
		num, _ := strconv.Atoi(s)
		nums[i] = num
	}

	fmt.Println("Number of arithmetic subsequences:", countArithmeticSlices(nums))
}
