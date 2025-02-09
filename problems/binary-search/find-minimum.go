package binarysearch

func FindMinimumBruteForce(nums []int) int {
	min := 99999999
	for _, val := range nums {
		if val < min {
			min = val
		}

	}
	return min

}

func FindMinOptimalApraochWithDuplicates(nums []int) int {
	size := len(nums)
	low := 0
	high := size - 1
	ans := 99999
	for low <= high {
		mid := (low + high) / 2
		ans = min(ans, nums[mid])

		if nums[low] == nums[mid] && nums[mid] == nums[high] {
			low = low + 1
			high = high - 1

		}

		if nums[low] <= nums[mid] {
			ans = min(ans, nums[low])
			low = mid + 1
		} else {
			ans = min(ans, nums[mid])
			high = mid - 1
		}
	}
	return ans

}
func FindMinOptimalApraoch(nums []int) int {
	size := len(nums)
	ans := 999999
	low := 0
	high := size - 1
	for low <= high {
		mid := (low + high) / 2
		// figured out mid
		// figure out which part is sorted take the minimum from it compare it against answer if smaller replace it and eliminate that half
		if nums[low] <= nums[high] {

		}
		if nums[low] <= nums[mid] { //left part is sorted
			if nums[low] < ans {
				ans = nums[low]
			}

			low = mid + 1
		} else {
			if nums[mid] < ans {
				ans = nums[mid]
			}

			high = mid - 1
		}
	}
	return ans
}
