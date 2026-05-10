package main

import (
	"fmt"

	"github.com/dsa/heap"
	"github.com/dsa/problems/implementations"
)

func main() {
	//input := []int{2, 4, 6, 8, 10, 12}
	//head := linkedlist.ArrToLinkedList(input)

	//	head = linkedlist.InsertAtEnd(head, 9)
	//	linkedlist.PrintLinkedList(head)
	//	x := linkedlist.LengthOfLinkedList(head)
	//	fmt.Println(x)
	//	head = deletenode.DeleteHead(head)
	//	linkedlist.PrintLinkedList(head)
	//	fmt.Println("deleting a tail of linked list")
	//	head = deletenode.DeleteTail(head)
	//	err := linkedlist.PrintLinkedList(head)
	//	if err != nil {
	//		println("linked list is empty")
	//	}
	//	NodeTobedeleted, err := linkedlist.Search(head, 10)
	//	if err != nil {
	//		fmt.Println("linked list is empty")

	//	}

	//	deletenode.DeleteNode(NodeTobedeleted)
	//
	// printing linked list after deelting 10 let's see
	//
	//	linkedlist.PrintLinkedList(head)
	//	strings.ReverseString("the sky is blue")
	//	area := stack.LargestRectangle([]int{2, 1, 5, 6, 2, 3})
	//	fmt.Println("area", area)
	//mock := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//ans := stack.MaxSlidingWindow(mock, 3)
	//fmt.Println(ans)
	//newNode.PrintLinkedList()a
	//
	lru := implementations.NewLruCacheWithCapacity(2)
	lru.Put(1, 1)
	lru.Put(2, 2)

	fmt.Println(lru.Get(1))

	lru.Put(3, 3)

	fmt.Println(lru.Get(2))

	lru.Put(4, 4)

	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))

	maxHeap := heap.NewMaxArrayHeap()
	maxHeap.Insert(50)
	maxHeap.Insert(55)
	maxHeap.Insert(53)
	maxHeap.Insert(52)
	maxHeap.Insert(54)
	maxHeap.Print()

	fmt.Println("AFTER DELETING ROOT NODE")

	maxHeap.Delete()
	maxHeap.Print()

	arr := []int{-1, 26, 23, 24, 30, 35, 40}

	for i := len(arr) / 2; i > 0; i-- {
		heap.Heapify(arr, len(arr), i)
	}

	for _, val := range arr {
		fmt.Println(val)
	}

}
