package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findDuplicate(nums []int) []int {
	var result []int

	for _, num := range nums {
		index := abs(num) - 1
		if nums[index] < 0 {
			result = append(result, abs(num))
		} else {
			nums[index] = -nums[index]
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter numbers separated by comma: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strNums := strings.Split(input, ",")

	nums := make([]int, len(strNums))
	for i, str := range strNums {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Invalid input: ", str)
			return
		}

		nums[i] = num
	}

	result := findDuplicate(nums)
	fmt.Println("Duplicate in the array: ", result)
}
