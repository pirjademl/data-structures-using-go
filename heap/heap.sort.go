package heap

import (
	"cmp"
	"fmt"
)

func SortArrayUsingHeapSort[T cmp.Ordered](nums []T) []T {
	n := len(nums)
	k := n/2 - 1

	for i := k; i >= 0; i-- {
		Heapify(nums, n, i)
	}
	fmt.Println("After heapify")
	for _, val := range nums {
		fmt.Println(val)

	}
	n = n - 1

	for n >= 1 {
		// swap
		nums[n], nums[0] = nums[0], nums[n]
		Heapify(nums, n, 0)
		n--
	}
	return nums

}
