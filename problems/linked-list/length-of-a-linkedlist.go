package linkedlist

func LengthOfLinkedList(head *Node) int {
	curr := head
	count := 0
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count

}
