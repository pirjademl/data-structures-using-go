package linkedlist

func InsertAtEnd(head *Node, x int) *Node {
	// I do have the access to the head
	// what if the linked list is empty in itself
	// if yes then make a new node and attach it to the head itself
	// otherwise go to the end where the node's next pointer is null and create a new node and assign the next of this curr pointer to the newly created node itself
	var curr *Node
	if head == nil {
		newNode := &Node{x, nil}
		head = newNode
		return head

	}
	curr = head
	for curr.Next != nil {
		curr = curr.Next
	}
	temp := &Node{x, nil}
	curr.Next = temp
	return head

}
