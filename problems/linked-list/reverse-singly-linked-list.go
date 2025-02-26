package linkedlist

func ReveserLinkedListIterative(head *Node) *Node {
	curr := head
	var prev *Node = nil
	for curr != nil {
		front := curr.Next
		curr.Next = prev
		prev = curr
		curr = front
	}
	return prev

}
func ReveserLinkedListRecursive(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReveserLinkedListRecursive(head.Next)
	front := head.Next
	front.Next = head
	head.Next = nil
	return newHead

}
