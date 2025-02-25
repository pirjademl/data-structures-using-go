package linkedlist

type DLLNode struct {
	Data int
	Prev *DLLNode
	Next *DLLNode
}

func ReverseLinkedList(head *DLLNode) *DLLNode {
	curr := head
	prev := curr
	for curr != nil {
		prev = curr.Prev
		curr.Prev = curr.Next
		curr.Next = prev
		curr = curr.Prev
	}
	return prev.Prev
}
