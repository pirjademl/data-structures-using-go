package binarysearch

func occurence(mat []int) int {
	size := len(mat)
	low := 0
	high := size - 1
	for low <= high {
		mid := (low + high) / 2
		if mat[mid] < 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low

}
func MaxRowWithCount1(mat [][]int) int {
	size := len(mat)
	row := -1
	count := 0
	for i := 0; i < size; i++ {
		index := occurence(mat[i])
		oneCount := size - index
		if oneCount > count {
			row = i
		}
	}
	return row

}
