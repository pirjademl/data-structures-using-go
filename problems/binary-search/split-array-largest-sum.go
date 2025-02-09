package binarysearch

import "slices"

func checkSum(nums []int, k int, min int) bool {
	count := 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+prev > min {
			count++
			prev = nums[i]
		} else {
			prev += nums[i]
		}
	}
	if count <= k {
		return true
	}
	return false
}

func LargestSplitSum(nums []int, k int) int {
	low := slices.Max(nums)
	high := sumOfArray(nums)
	for i := low; i <= high; i++ {
		if checkSum(nums, k, i) {
			return i
		}
	}
	return -1
}
