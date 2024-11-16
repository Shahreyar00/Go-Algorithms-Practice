package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	convertToMinutes := func(time string) int {
		parts := strings.Split(time, ":")
		hours, _ := strconv.Atoi(parts[0])
		minutes, _ := strconv.Atoi(parts[1])
		return hours*60 + minutes
	}

	minutes := make([]int, len(timePoints))
	for i, time := range timePoints {
		minutes[i] = convertToMinutes(time)
	}

	sort.Ints(minutes)
	minDiff := math.MaxInt32
	for i := 1; i < len(minutes); i++ {
		minDiff = min(minDiff, minutes[i]-minutes[i-1])
	}

	circularDiff := (1440 - minutes[len(minutes)-1]) + minutes[0]
	minDiff = min(minDiff, circularDiff)

	return minDiff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Println("Enter the number of time points:")
	fmt.Scan(&n)

	timePoints := make([]string, n)
	fmt.Println("Enter the time points in HH:MM format:")
	for i := 0; i < n; i++ {
		fmt.Scan(&timePoints[i])
	}

	result := findMinDifference(timePoints)
	fmt.Printf("The minimum minutes difference is: %d\n", result)
}
