package binarysearch

import "fmt"

func MedianOfTwoSortedArrayBruteForce(arr1 []int, arr2 []int) float64 {
	totalSize := len(arr1) + len(arr2)
	ans := make([]int, totalSize)
	index := 0
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			ans[index] = arr1[i]
			index++
			i++
		} else {
			ans[index] = arr2[j]
			j++
		}
	}
	for i < len(arr1) {
		ans[index] = arr1[i]
		index++
		i++
	}
	for i < len(arr2) {
		ans[index] = arr2[j]
		index++
		j++
	}
	if totalSize%2 != 0 {
		return float64(ans[totalSize/2])
	}
	median := (float64(ans[totalSize/2]) + float64(ans[(totalSize/2)-1])) / 2.00
	return median
}

func MedianOfTwoSortedArrayBetterAproach(arr1 []int, arr2 []int) float64 {
	totalSize := len(arr1) + len(arr2)
	index2 := totalSize / 2 // 2
	element2 := 0
	element1 := 0
	index1 := index2 - 1 // 1
	i := 0
	j := 0
	count := 0
	for i < len(arr1) && (j < len(arr2)) {
		if arr1[i] <= arr2[j] {
			if count == index1 {
				element1 = arr1[i]
			}
			if count == index2 {
				element2 = arr1[i]
			}
			i++
			count++
		} else {
			if count == index1 {
				element1 = arr2[j]
			}
			if count == index2 {
				element2 = arr2[j]
			}
			j++
			count++
		}
	}
	for i < len(arr1) {
		if count == index1 {
			element1 = arr1[i]
		}
		if count == index2 {

			element2 = arr1[i]
		}
		i++
		count++
	}
	for j < len(arr2) {
		if count == index1 {
			element1 = arr2[j]
		}
		if count == index2 {
			element2 = arr2[j]
		}
		j++
		count++
	}
	if ifodd := (totalSize % 2); ifodd != 0 {
		fmt.Println("size is odd")
		return float64(element2)
	}
	fmt.Println(element1, element2)
	median := (float64(element1) + float64(element2)) / 2.00
	return median
}
func MedianOfTwoSortedArrayOptimal(arr1 []int, arr2 []int) float64 {
	return 2.00
}
