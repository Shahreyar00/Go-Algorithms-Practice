package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	var leftmost int

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == 0 {
				leftmost = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return leftmost
}

func createTree(values []string) *TreeNode {
	if len(values) == 0 || values[0] == "null" {
		return nil
	}

	root := &TreeNode{}
	queue := []*TreeNode{root}
	i := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		i++
		if i < len(values) && values[i] != "null" {
			node.Left = &TreeNode{Val: parseValue(values[i])}
			queue = append(queue, node.Left)
		}

		i++
		if i < len(values) && values[i] != "null" {
			node.Right = &TreeNode{Val: parseValue(values[i])}
			queue = append(queue, node.Right)
		}
	}

	return root
}

func parseValue(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter tree nodes in level-order format (e.g., 1,2,3,4,null,5,6,null,null,7): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	values := strings.Split(input, ",")
	root := createTree(values)
	fmt.Println("Bottom-left value:", findBottomLeftValue(root))
}
