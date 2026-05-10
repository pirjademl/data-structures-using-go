package heap

import "fmt"

type MaxArrayHeap struct {
	size int
	arr  []int
}

func NewMaxArrayHeap() *MaxArrayHeap {
	return &MaxArrayHeap{
		size: 0,
		arr:  make([]int, 100),
	}
}
func (maxHeap *MaxArrayHeap) Insert(val int) {
	maxHeap.size++
	maxHeap.arr[maxHeap.size] = val
	i := maxHeap.size

	for i > 1 {
		parentI := i / 2
		if maxHeap.arr[parentI] < maxHeap.arr[i] {
			maxHeap.arr[parentI], maxHeap.arr[i] = maxHeap.arr[i], maxHeap.arr[parentI]
			i = parentI
		} else {
			return
		}

	}
}

func (maxHeap *MaxArrayHeap) Delete() {
	maxHeap.arr[1] = maxHeap.arr[maxHeap.size]
	maxHeap.size--
	root := 1

	for root < maxHeap.size {
		leftChild := 2 * root
		rightChild := 2*root + 1
		if leftChild < maxHeap.size && maxHeap.arr[root] < maxHeap.arr[leftChild] {
			maxHeap.arr[root], maxHeap.arr[leftChild] = maxHeap.arr[leftChild], maxHeap.arr[root]
			root = leftChild

		} else if rightChild < maxHeap.size && maxHeap.arr[root] < maxHeap.arr[rightChild] {

			maxHeap.arr[root], maxHeap.arr[rightChild] = maxHeap.arr[rightChild], maxHeap.arr[root]
			root = rightChild

		} else {
			return
		}

	}

}
func (heap *MaxArrayHeap) Print() {
	for i := 1; i <= heap.size; i++ {
		fmt.Println(heap.arr[i])
	}

}
