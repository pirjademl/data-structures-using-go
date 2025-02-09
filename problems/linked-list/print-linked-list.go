package linkedlist

import (
	"errors"
	"fmt"
)

func PrintLinkedList(head *Node) error {
	if head == nil {
		return errors.New("Linked list is empty")
	}
	curr := head
	for curr != nil {
		fmt.Println(curr.Data)
		curr = curr.Next
	}
	return nil

}
