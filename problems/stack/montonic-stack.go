package stack

func NextSmaller(arr []int) []int {
	n := len(arr)
	st := []int{}
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) != 0 && arr[i] <= st[len(st)-1] {
			st = st[:len(st)-1]
		}

		if len(st) != 0 {
			ans[i] = st[len(st)-1]
		} else {
			ans[i] = n
		}

		st = append(st, i)

	}
	return ans

}
func PrevSmaller(arr []int) []int {
	n := len(arr)
	st := []int{}
	ans := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) != 0 && arr[i] < st[len(st)-1] {
			st = st[:len(st)-1]
		}

		if len(st) != 0 {
			ans[i] = st[len(st)-1]
		} else {
			ans[i] = -1
		}
		st = append(st, i)

	}
	return ans
}
func LargestRectangle(heights []int) int {
	area := 0
	nsmalller := NextSmaller(heights)
	psmaller := PrevSmaller(heights)

	for i, val := range heights {
		right := nsmalller[i]
		left := psmaller[i]
		local := ((right - left) - 1) * val
		if local > area {
			area = local
		}
	}
	return area
}
