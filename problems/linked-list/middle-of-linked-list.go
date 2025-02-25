package linkedlist

import "math"

func middleOfLinkedList(head *Node) *Node {
	curr := head
	total := 0
	for curr != nil {
		total++
		curr = curr.Next
	}
	middle := 0
	if total%2 != 0 {
		middle = int(math.Ceil(float64(total) / 2.0))
	} else {
		middle = int(math.Ceil(float64(total)/2.0) + 1)
	}
	index := 1
	curr = head
	for curr != nil {
		if index == middle {
			head = curr
			break
		}
		curr = curr.Next
		index++

	}
	return head

}
