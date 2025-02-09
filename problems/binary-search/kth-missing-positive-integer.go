package binarysearch

import (
	"slices"
)

func binarySearch(nums []int, target int) bool {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func KMissingPositiveNumber(nums []int, k int) int {
	max := slices.Max(nums)
	count := 0
	for i := 1; i <= max; i++ {
		found := binarySearch(nums, i)
		if !found {
			count++
		}
		if count == k {
			return i
		}
	}
	return -1
}
