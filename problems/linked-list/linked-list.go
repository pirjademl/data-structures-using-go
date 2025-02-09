package linkedlist

type Node struct {
	Data int
	Next *Node
}
type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtFront(data int) {
	newNode := &Node{data, nil}
	if list.head == nil {
		list.head = newNode
	}
	curr := list.head
	for curr.Next != nil {

	}
}
