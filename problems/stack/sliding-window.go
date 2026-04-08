package stack

func MaxSlidingWindow(nums []int, k int) []int {
	size := (len(nums) - k) + 1
	ans := make([]int, 0, size)
	deq := []int{}
	for i, val := range nums {
		if len(deq) != 0 && deq[0] <= (i-k) {
			deq = deq[1:]
		}
		for len(deq) != 0 && nums[deq[len(deq)-1]] <= val {
			deq = deq[:len(deq)-1]
		}

		deq = append(deq, i)
		if i >= k-1 {
			ans = append(ans, nums[deq[0]])
		}

	}
	return ans
}
