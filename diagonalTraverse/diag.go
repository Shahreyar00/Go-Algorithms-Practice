package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	result := make([]int, 0, m*n)

	for d := 0; d < m+n-1; d++ {
		intermediate := []int{}
		var r, c int
		if d < n {
			r, c = 0, d
		} else {
			r, c = d-n+1, n-1
		}

		for r < m && c >= 0 {
			intermediate = append(intermediate, mat[r][c])
			r++
			c--
		}

		if d%2 == 0 {
			for i := len(intermediate) - 1; i >= 0; i-- {
				result = append(result, intermediate[i])
			}
		} else {
			result = append(result, intermediate...)
		}
	}
	return result
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter matrix rows, one per line (space-separated integers). Enter a blank line to finish:")
	matrix := [][]int{}

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		row := []int{}
		for _, numStr := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(numStr)
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	result := findDiagonalOrder(matrix)
	fmt.Println("Diagonal order of the matrix:", result)
}
