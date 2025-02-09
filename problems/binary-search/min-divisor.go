package binarysearch

import "math"

func sumOfDivisor(nums []int, divisor int) int {
	sum := 0
	for _, val := range nums {
		sum += int(math.Ceil(float64(val) / float64(divisor)))
		println("sum is ", sum)
	}
	return sum
}

func MinDivisor(nums []int, threshold int) int {

	for i := 1; i <= threshold; i++ {
		sum := sumOfDivisor(nums, i)
		if sum <= threshold {
			return i
		}
	}
	return -1
}
