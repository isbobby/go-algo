package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisjointSetNaive(t *testing.T) {
	size := 5

	dsu := InitializeDisjointSet(size)

	for i := 0; i < size; i++ {
		dsu.MakeSet(i)
	}

	dsu.NaiveUnion(0, 2)
	dsu.NaiveUnion(1, 3)
	dsu.NaiveUnion(0, 4)

	assert.True(t, dsu.NaiveSameSet(0, 2))
	assert.True(t, dsu.NaiveSameSet(2, 4))
	assert.True(t, dsu.NaiveSameSet(1, 3))
	assert.False(t, dsu.NaiveSameSet(0, 3))
	assert.False(t, dsu.NaiveSameSet(3, 4))
}
