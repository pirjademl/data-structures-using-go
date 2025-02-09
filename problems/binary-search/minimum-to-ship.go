package binarysearch

import "slices"

func sumOfArray(nums []int) (sum int) {
	for _, val := range nums {
		sum += val

	}
	return sum
}

func daysToRequire(nums []int, capacity int) int {
	days, load := 1, 0
	for _, val := range nums {
		if load+val > capacity {
			days++
			load = val
		} else {
			load += val
		}
	}
	return days
}

func MinToShip(nums []int, days int) int {
	low := slices.Max(nums)
	high := sumOfArray(nums)
	for low <= high {
		mid := (low + high) / 2
		if val := daysToRequire(nums, mid); val <= days {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
