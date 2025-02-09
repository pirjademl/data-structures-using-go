package binarysearch

func SearchInRotatedSortedArray(nums []int, target int) int {
	size := len(nums)
	low := 0
	high := size - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[low] <= nums[mid] { // left part is sorted
			if target >= nums[low] && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}

		} else { //right part is sorted
			if target >= nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}

		}

	}
	return -1
}
