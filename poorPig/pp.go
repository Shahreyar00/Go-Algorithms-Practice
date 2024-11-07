package main

import (
	"fmt"
	"math"
)

func minPigs(buckets, minutesToDie, minutesToTest int) int {
	tests := minutesToTest / minutesToDie
	pigs := 0

	for math.Pow(float64(tests+1), float64(pigs)) < float64(buckets) {
		pigs++
	}

	return pigs
}

func main() {
	var buckets, minutesToDie, minutesToTest int

	fmt.Println("Enter the number of buckets:")
	fmt.Scan(&buckets)
	fmt.Println("Enter the minutesToDie:")
	fmt.Scan(&minutesToDie)
	fmt.Println("Enter the minutesToTest:")
	fmt.Scan(&minutesToTest)

	result := minPigs(buckets, minutesToDie, minutesToTest)
	fmt.Printf("Minimum number of pigs needed: %d\n", result)
}
