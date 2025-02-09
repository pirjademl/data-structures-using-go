package binarysearch

import "slices"

func possibleAnswer(bloomdays []int, day int, bouquets int, flowers int) bool {
	counter := 0
	answer := 0
	for _, val := range bloomdays {
		if val <= day {
			counter++
		} else {
			answer += (counter / flowers)
		}
	}
	answer += (counter / flowers)
	if answer >= bouquets {
		return true
	}
	return false

}
func MinimumdaysToMakeMbouqets(bloomdays []int, m, k int) int {
	if len(bloomdays) < m*k {
		return -1
	}

	low := slices.Min(bloomdays)
	high := slices.Max(bloomdays)
	for low <= high {
		mid := (low + high) / 2
		if possibleAnswer(bloomdays, mid, m, k) {
			high = mid - 1

		} else {
			low = mid + 1
		}

	}
	return low
}
