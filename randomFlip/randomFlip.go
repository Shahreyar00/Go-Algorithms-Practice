package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	m, n      int
	total     int
	flipped   map[int]int
	available int
}

func Constructor(m int, n int) Solution {
	return Solution{
		m:         m,
		n:         n,
		total:     m * n,
		flipped:   make(map[int]int),
		available: m * n,
	}
}

func (this *Solution) Flip() []int {
	randIndex := rand.Intn(this.available)
	cell := this.flipped[randIndex]
	if cell == 0 {
		cell = randIndex
	}

	this.available--

	this.flipped[randIndex] = this.flipped[this.available]
	if this.flipped[randIndex] == 0 {
		this.flipped[randIndex] = this.available
	}

	return []int{cell / this.n, cell % this.n}
}

func (this *Solution) Reset() {
	this.flipped = make(map[int]int)
	this.available = this.total
}

func main() {
	obj := Constructor(3, 1)
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	obj.Reset()
	fmt.Println(obj.Flip())
}
