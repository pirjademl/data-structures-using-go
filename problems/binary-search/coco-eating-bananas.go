package binarysearch

func maxiumElementInArray(p []int) int {
	maxi := -1
	for _, val := range p {
		if val > maxi {
			maxi = val
		}

	}
	return maxi

}
func TotalEatingTimea(piles []int, hourly int) int {
	totalTime := 0
	for _, val := range piles {
		totalTime = totalTime + (val+hourly-1)/hourly

	}
	return totalTime
}

func MinimumCount(piles []int, h int) int {
	//
	//
	low := 1
	high := maxiumElementInArray(piles)

	for low <= high {
		mid := (low + high) / 2
		// calculate how much time it takes if the speed is mid
		time := TotalEatingTimea(piles, mid)
		if time <= h {
			high = mid - 1

		} else {
			low = mid + 1
		}

	}
	return low

}
