package main

import (
	"fmt"

	linkedlist "github.com/dsa/problems/linked-list"
	deletenode "github.com/dsa/problems/linked-list/delete-node"
)

func main() {
	input := []int{2, 4, 6, 8, 10, 12}
	head := linkedlist.ArrToLinkedList(input)

	//	head = linkedlist.InsertAtEnd(head, 9)
	linkedlist.PrintLinkedList(head)
	x := linkedlist.LengthOfLinkedList(head)
	fmt.Println(x)
	//	head = deletenode.DeleteHead(head)
	//	linkedlist.PrintLinkedList(head)
	//	fmt.Println("deleting a tail of linked list")
	//	head = deletenode.DeleteTail(head)
	//	err := linkedlist.PrintLinkedList(head)
	//	if err != nil {
	//		println("linked list is empty")
	//	}
	NodeTobedeleted, err := linkedlist.Search(head, 10)
	if err != nil {
		fmt.Println("linked list is empty")

	}

	deletenode.DeleteNode(NodeTobedeleted)
	// printing linked list after deelting 10 let's see
	linkedlist.PrintLinkedList(head)

}
