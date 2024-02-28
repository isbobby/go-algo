package compaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompactSegment(t *testing.T) {
	rowsA := []row{
		{1, 1}, {1, 2},
	}
	rowsB := []row{
		{2, 2}, {3, 1}, {1, 4},
	}
	expectedRows := []row{
		{2, 2}, {3, 1}, {1, 4},
	}
	gotSegment := compactSegments(&segment{rowsA}, &segment{rowsB})
	assert.Equal(t, expectedRows, gotSegment.rows)
}
