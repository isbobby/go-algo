package mst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestL1584Case1(t *testing.T) {
	coordinates := [][]int{
		{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
	}

	expected := 20

	got := minCostConnectPoints(coordinates)

	assert.Equal(t, expected, got)

	coordinates2 := [][]int{
		{0, 0}, {1, 1}, {1, 0}, {-1, 1},
	}

	got2 := minCostConnectPoints(coordinates2)

	assert.Equal(t, 4, got2)
}

func TestL1584Case5(t *testing.T) {
	coordinates := [][]int{
		{0, 0}, {1, 1}, {1, 0}, {-1, 1},
	}

	got := minCostConnectPoints(coordinates)

	assert.Equal(t, 4, got)
}
