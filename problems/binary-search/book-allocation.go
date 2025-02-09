package binarysearch

import "slices"

func isPossible(books []int, min int) int {
	count := 1
	pages := books[0]
	for i := 1; i < len(books); i++ {
		if books[i]+pages > min {
			count++
			pages = books[i]
		} else {
			pages += books[i]
		}
	}
	return count
}

func MinPages(books []int, students int) int {
	//one observation is that the answer is lies between high and the sum of array elements
	low := slices.Max(books)
	high := sumOfArray(books)
	for low <= high {
		mid := (low + high) / 2
		if ans := isPossible(books, mid); ans > students {
			low = mid + 1

		} else {
			high = mid - 1
		}
	}
	return low
}
