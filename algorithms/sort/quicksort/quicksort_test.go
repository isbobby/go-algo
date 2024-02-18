package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortIntArry(t *testing.T) {
	//			   4, 6, 3, 8, 7, 1, 2
	input := []int{4, 6, 1, 3, 8, 7, 2}

	QuicksortIntArray(input)

	expected := []int{1, 2, 3, 4, 6, 7, 8}

	assert.Equal(t, expected, input, "input:{4, 6, 1, 3, 8, 7, 2}")

	input2 := []int{5, 2, 3, 1}

	QuicksortIntArray(input2)

	expected2 := []int{1, 2, 3, 5}

	assert.Equal(t, expected2, input2, "input:{5, 2, 3, 1}")
}
