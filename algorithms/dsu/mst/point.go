package mst

import "math"

type point struct {
	x int
	y int
}

func pointDistance(a, b point) float64 {
	x := a.x - b.x
	y := a.y - b.y
	return math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
}

func manhattanDistance(a, b point) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func (p point) equal(a point) bool {
	return p.x == a.x && p.y == a.y
}
