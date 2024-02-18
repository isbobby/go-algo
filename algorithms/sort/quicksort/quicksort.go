package quicksort

import "fmt"

func QuicksortIntArray(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) {
	// only sort if the boundary specified needs to be sorted
	if low < high {
		fmt.Printf("index:(%v,%v)\n", low, high)
		pivotIndex := partition(arr, low, high)

		// left sub-array
		quickSort(arr, low, pivotIndex-1)

		// right sub-array
		quickSort(arr, pivotIndex+1, high)
	}
}

// partition with the last element as the pivot, after
// partitioning, the pviot will be in the sorted index
//
// this sorted index is returned
func partition(arr []int, low, high int) int {
	pivot := arr[high]

	l := low - 1

	for r := low; r < high; r++ {
		if arr[r] < pivot {
			l++

			arr[r], arr[l] = arr[l], arr[r]
		}
	}

	arr[high], arr[l+1] = arr[l+1], arr[high]

	return l + 1
}
