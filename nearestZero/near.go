package main

import (
	"container/list"
	"fmt"
)

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		for j := range result[i] {
			if mat[i][j] == 1 {
				result[i][j] = 1e9
			}
		}
	}

	q := list.New()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				q.PushBack([2]int{i, j})
			}
		}
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for q.Len() > 0 {
		cell := q.Remove(q.Front()).([2]int)
		x, y := cell[0], cell[1]

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]

			if nx >= 0 && nx < m && ny >= 0 && ny < n && result[nx][ny] > result[x][y]+1 {
				result[nx][ny] = result[x][y] + 1
				q.PushBack([2]int{nx, ny})
			}
		}
	}

	return result
}

func main() {
	var m, n int
	fmt.Print("Enter number of rows (m): ")
	fmt.Scan(&m)
	fmt.Print("Enter number of columns (n): ")
	fmt.Scan(&n)

	fmt.Println("Enter the binary matrix row by row (space-separated):")
	mat := make([][]int, m)
	for i := range mat {
		mat[i] = make([]int, n)
		for j := range mat[i] {
			fmt.Scan(&mat[i][j])
		}
	}

	result := updateMatrix(mat)
	fmt.Println("Output matrix with distances:")
	for _, row := range result {
		fmt.Println(row)
	}
}
