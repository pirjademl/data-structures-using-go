package binarysearch

func firstOccurence(arr []int, target int) int {
	first := -1
	size := len(arr)
	low := 0
	high := size - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			first = mid
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return first
}

func lastOccurence(arr []int, target int) int {
	last := -1
	size := len(arr)
	low := 0
	high := size - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			last = mid
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return last
}

// Count Occurence is a function which returns the occurence of  a element in a sorted array
// @arg 1 arr
// @arg 2 target element
func CountOccurence(arr []int, target int) int {
	first := firstOccurence(arr, target)
	if first == -1 {
		return 0
	}
	last := lastOccurence(arr, target)
	return last - first + 1
}
