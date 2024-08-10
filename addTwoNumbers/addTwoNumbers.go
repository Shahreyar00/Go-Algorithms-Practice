package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func pushToStack(head *ListNode, stack *[]int) {
	for head != nil {
		*stack = append(*stack, head.Val)
		head = head.Next
	}
}

func createLinkedListFromStack(stack *[]int) *ListNode {
	var head *ListNode
	for len(*stack) > 0 {
		val := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		newNode := &ListNode{Val: val}
		newNode.Next = head
		head = newNode
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := []int{}
	stack2 := []int{}

	pushToStack(l1, &stack1)
	pushToStack(l2, &stack2)

	var carry int
	var result *ListNode

	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		sum := carry
		if len(stack1) > 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		carry = sum / 10
		node := &ListNode{Val: sum % 10}
		node.Next = result
		result = node
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first linked list (comma-separated): ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	nums1 := strings.Split(input1, ",")
	var l1, curr1 *ListNode

	for _, num := range nums1 {
		val, _ := strconv.Atoi(num)
		if l1 == nil {
			l1 = &ListNode{Val: val}
			curr1 = l1
		} else {
			curr1.Next = &ListNode{Val: val}
			curr1 = curr1.Next
		}
	}

	fmt.Print("Enter second linked list (comma-separated): ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	nums2 := strings.Split(input2, ",")
	var l2, curr2 *ListNode

	for _, num := range nums2 {
		val, _ := strconv.Atoi(num)
		if l2 == nil {
			l2 = &ListNode{Val: val}
			curr2 = l2
		} else {
			curr2.Next = &ListNode{Val: val}
			curr2 = curr2.Next
		}
	}

	result := addTwoNumbers(l1, l2)

	fmt.Print("Result linked list: ")
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
		if result != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
