package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findPattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	stack := []int{}
	third := -1 << 31

	for i := n - 1; i >= 0; i-- {
		if nums[i] < third {
			return true
		}

		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			third = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, nums[i])
	}

	return false
}

func main() {
	fmt.Println("Enter the array elements separated by spaces: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	numStrings := strings.Split(input, " ")
	nums := []int{}
	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter integers only.")
			return
		}
		nums = append(nums, num)
	}

	if findPattern(nums) {
		fmt.Println("132 pattern found.")
	} else {
		fmt.Println("132 pattern NOT found.")
	}
}
