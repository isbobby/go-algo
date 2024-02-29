package topology

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopologicalSort(t *testing.T) {
	nodes := []int{1, 2, 3, 4, 5}
	edges := map[int][]int{
		1: {3},
		2: {1},
		3: {5},
		5: {4},
	}
	expects := []int{2, 1, 3, 5, 4}

	got := TopologicalSort(edges, nodes)

	assert.Equal(t, expects, got)
}
