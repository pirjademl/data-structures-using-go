package linkedlist

import "errors"

func Search(head *Node, x int) (*Node, error) {
	if head == nil {
		return nil, errors.New("linked list is empty")
	}
	curr := head
	for curr != nil {
		if curr.Data == x {
			return curr, nil
		}
		curr = curr.Next
	}
	return nil, nil
}
