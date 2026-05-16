package heap

import "cmp"

func Heapify[T cmp.Ordered](arr []T, n, i int) {
	largest := i

	leftIndex := 2*largest + 1
	rightIndex := 2*largest + 2

	if leftIndex < n && arr[largest] <= arr[leftIndex] {
		largest = leftIndex

	}
	if rightIndex < n && arr[largest] <= arr[rightIndex] {
		largest = rightIndex
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		Heapify(arr, n, largest)
	}

}
