package binarysearch

import "slices"

func isMinimum(nums []int, k int, min int) bool {
	count := 1
	prev := nums[0]
	for _, val := range nums {
		if val+prev > min {
			count++
			prev = val
		} else {
			prev += val
		}
	}
	if count <= k {
		return true
	}
	return false
}

func PainterPartition(nums []int, k int) int {
	low := slices.Max(nums)
	high := sumOfArray(nums)
	for i := low; i <= high; i++ {
		if isMinimum(nums, k, i) {
			return i
		}

	}
	return -1

}
