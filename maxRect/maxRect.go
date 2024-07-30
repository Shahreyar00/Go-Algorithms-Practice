package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	maxSum := int(^uint(0)>>1) * -1

	for left := 0; left < n; left++ {
		rowSum := make([]int, m)
		for right := left; right < n; right++ {
			for i := 0; i < m; i++ {
				rowSum[i] += matrix[i][right]
			}
			maxSum = max(maxSum, maxSumNoLargerThanK(rowSum, k))
		}
	}

	return maxSum
}

func maxSumNoLargerThanK(arr []int, k int) int {
	maxSum := int(^uint(0)>>1) * -1
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum <= k {
				maxSum = max(maxSum, sum)
			}
		}
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the number of rows:")
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter the number of columns:")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	matrix := make([][]int, m)
	fmt.Println("Enter the matrix row by row (space separated values):")
	for i := 0; i < m; i++ {
		scanner.Scan()
		row := strings.Split(scanner.Text(), " ")
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j], _ = strconv.Atoi(row[j])
		}
	}

	fmt.Println("Enter the value of k:")
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	result := maxSumSubmatrix(matrix, k)
	fmt.Printf("The max sum of a rectangle no larger than %d is: %d\n", k, result)
}
