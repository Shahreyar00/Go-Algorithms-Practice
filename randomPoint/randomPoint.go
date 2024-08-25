package main

import (
	"math/rand"
	"sort"
	"time"
)

type Solution struct {
	rects       [][]int
	pointsSum   []int
	totalPoints int
}

func Constructor(rects [][]int) Solution {
	s := Solution{
		rects: rects,
	}

	rand.Seed(time.Now().UnixNano())

	s.pointsSum = make([]int, len(rects))
	for i, rect := range rects {
		width := rect[2] - rect[0] + 1
		height := rect[3] - rect[1] + 1
		points := width * height
		if i > 0 {
			s.pointsSum[i] = s.pointsSum[i-1] + points
		} else {
			s.pointsSum[i] = points
		}
	}

	s.totalPoints = s.pointsSum[len(rects)-1]

	return s
}

func (s *Solution) Pick() []int {
	pointIndex := rand.Intn(s.totalPoints)

	rectIndex := sort.Search(len(s.pointsSum), func(i int) bool {
		return s.pointsSum[i] > pointIndex
	})

	rect := s.rects[rectIndex]

	prevSum := 0
	if rectIndex > 0 {
		prevSum = s.pointsSum[rectIndex-1]
	}

	offset := pointIndex - prevSum
	width := rect[2] - rect[0] + 1

	row := offset / width
	col := offset % width

	return []int{rect[0] + col, rect[1] + row}
}

func main() {
	rects := [][]int{{-2, -2, 1, 1}, {2, 2, 4, 6}}
	obj := Constructor(rects)

	for i := 0; i < 5; i++ {
		point := obj.Pick()
		println(point[0], point[1])
	}
}
