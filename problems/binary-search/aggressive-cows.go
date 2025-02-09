package binarysearch

func check(stalls []int, k int, minimum int) bool {
	count := 1
	prev := stalls[0]
	for _, _val := range stalls {
		if _val+prev > minimum {
			count++
		} else {
			prev += _val
		}
	}
	if count >= k {
		return true
	}
	return false
}

func AggressiveCows(stalls []int, k int) int {
	low := 0
	high := stalls[len(stalls)-1] - stalls[0]
	for low <= high {
		mid := (low + high) / 2
		if check(stalls, k, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return high
}
