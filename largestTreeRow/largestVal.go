package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Function to find the largest value in each tree row
func largestValues(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    var result []int
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        levelSize := len(queue)
        maxVal := queue[0].Val

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]

            if node.Val > maxVal {
                maxVal = node.Val
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, maxVal)
    }

    return result
}

// Function to build the tree from input values
func buildTree(values []string) *TreeNode {
    if len(values) == 0 || values[0] == "null" {
        return nil
    }

    rootVal, _ := strconv.Atoi(values[0])
    root := &TreeNode{Val: rootVal}
    queue := []*TreeNode{root}

    index := 1
    for index < len(values) {
        node := queue[0]
        queue = queue[1:]

        if index < len(values) && values[index] != "null" {
            leftVal, _ := strconv.Atoi(values[index])
            node.Left = &TreeNode{Val: leftVal}
            queue = append(queue, node.Left)
        }
        index++

        if index < len(values) && values[index] != "null" {
            rightVal, _ := strconv.Atoi(values[index])
            node.Right = &TreeNode{Val: rightVal}
            queue = append(queue, node.Right)
        }
        index++
    }

    return root
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Enter the tree values (comma-separated, use 'null' for empty nodes):")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    values := strings.Split(input, ",")

    root := buildTree(values)
    result := largestValues(root)

    fmt.Println("Largest values in each tree row:", result)
}
