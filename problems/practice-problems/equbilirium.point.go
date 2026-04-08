package practiceproblems

import "fmt"

var a int

func Eq(arr []int) {
	len := len(arr)
	index := 0
	for i := 0; i < len; i++ {

		leftsum := 0
		rightsum := 0
		if i > 0 {
			for j := i + 1; j < len; j++ {
				leftsum += arr[j]
			}

		}

		if i > 0 {
			for k := 0; k < i; k++ {
				rightsum += arr[k]
			}
		}
		if leftsum == rightsum && leftsum != 0 && rightsum != 0 {
			fmt.Print("sum is found equal", leftsum, rightsum)
			index = i
		}
	}
	println("this is the index ", index)

}
