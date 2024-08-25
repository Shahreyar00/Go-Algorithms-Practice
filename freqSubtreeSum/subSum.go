package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func calculateSubtreeSums(node *TreeNode, freqMap map[int]int) int {
	if node == nil {
		return 0
	}

	leftSum := calculateSubtreeSums(node.Left, freqMap)
	rightSum := calculateSubtreeSums(node.Right, freqMap)
	subtreeSum := node.Val + leftSum + rightSum

	freqMap[subtreeSum]++
	return subtreeSum
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	freqMap := make(map[int]int)
	calculateSubtreeSums(root, freqMap)

	maxFreq := 0
	for _, freq := range freqMap {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	var result []int
	for sum, freq := range freqMap {
		if freq == maxFreq {
			result = append(result, sum)
		}
	}

	return result
}

func buildTree(input []int, index int) *TreeNode {
	if index >= len(input) || input[index] == -1 {
		return nil
	}

	root := &TreeNode{Val: input[index]}
	root.Left = buildTree(input, 2*index+1)
	root.Right = buildTree(input, 2*index+2)
	return root
}

func main() {
	var n int
	fmt.Println("Enter number of elements in the tree (use -1 for null nodes):")
	fmt.Scan(&n)

	input := make([]int, n)
	fmt.Println("Enter the elements in lever-order format:")
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}

	root := buildTree(input, 0)
	result := findFrequentTreeSum(root)
	fmt.Println("Most frequent subtree sum(s):", result)
}
