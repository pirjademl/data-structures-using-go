package linkedlist

import "math"

func middleOfLinkedListBruteForce(head *Node) *Node {
	if head == nil {
		return nil
	}
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
func MiddleOfLinkedListTortoiseMethod(head *Node) *Node {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
