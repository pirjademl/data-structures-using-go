package binarysearch

func SingleElement(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[size-1] != nums[size-2] {
		return nums[size-1]
	}
	low := 1
	high := size - 2

	for low <= high {
		mid := (low + high) / 2
		//check if the mid is the answer itself
		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}

		//where am i on the right or left
		if mid%2 != 0 &&
			mid-1%2 == 0 { //i am on the left hand side and need to eliminate the left hand side
			low = mid + 1

		} else {
			high = mid - 1
		}
	}
	return -1
}
