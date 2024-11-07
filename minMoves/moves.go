package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minMoves(nums []int) int {
	sum, min := 0, nums[0]
	for _, num := range nums {
		sum += num
		if num < min {
			min = num
		}
	}

	return sum - (min * len(nums))
}

func main() {
	fmt.Println("Enter the array elements separated by spaces:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	numStrings := strings.Split(input, " ")
	nums := []int{}
	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter only integers.")
			return
		}
		nums = append(nums, num)
	}

	result := minMoves(nums)
	fmt.Printf("Minimum moves required: %d\n", result)
}
