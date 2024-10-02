package main

import (
	"fmt"
)

var directions = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1},
}

func updateBoard(board [][]byte, click []int) [][]byte {
	row, col := click[0], click[1]

	// If the clicked cell is a mine, game over, mark it as 'X'
	if board[row][col] == 'M' {
		board[row][col] = 'X'
		return board
	}

	// Otherwise, perform DFS to reveal cells
	dfs(board, row, col)
	return board
}

func dfs(board [][]byte, row, col int) {
	// If the cell is out of bounds or already revealed, return
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] != 'E' {
		return
	}

	mineCount := countAdjacentMines(board, row, col)

	if mineCount > 0 {
		// If there are adjacent mines, reveal the count
		board[row][col] = byte(mineCount + '0') // Convert the count to a char '1'-'8'
	} else {
		// If no adjacent mines, mark the cell as 'B' and continue revealing
		board[row][col] = 'B'
		// Recursively reveal adjacent cells
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			dfs(board, newRow, newCol)
		}
	}
}

// Helper function to count adjacent mines
func countAdjacentMines(board [][]byte, row, col int) int {
	count := 0
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow >= 0 && newRow < len(board) && newCol >= 0 && newCol < len(board[0]) && board[newRow][newCol] == 'M' {
			count++
		}
	}
	return count
}

func main() {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}

	var click = []int{3, 0}

	updatedBoard := updateBoard(board, click)
	for _, row := range updatedBoard {
		fmt.Println(string(row))
	}
}
