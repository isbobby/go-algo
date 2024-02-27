package mst

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	pointA := point{
		x: 0, y: 0,
	}

	pointB := point{
		x: 2, y: 2,
	}

	got := pointDistance(pointA, pointB)

	assert.Equal(t, math.Sqrt(8), got)
}

func TestManhattanDistance(t *testing.T) {
	pointA := point{
		x: 0, y: 0,
	}

	pointB := point{
		x: 2, y: 2,
	}

	got := manhattanDistance(pointA, pointB)

	assert.Equal(t, 4, got)

	got2 := manhattanDistance(pointB, pointA)

	assert.Equal(t, 4, got2)
}
