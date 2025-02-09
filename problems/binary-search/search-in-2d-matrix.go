package binarysearch

func SearchIn2dMatrix(nums [][]int, target int) bool {
	rowSize := len(nums)
	colSize := len(nums[0])
	for i := 0; i < rowSize; i++ {
		if nums[i][0] <= target && target <= nums[i][colSize-1] {
			return binarySearch(nums[i], target)

		}
	}
	return false

}
