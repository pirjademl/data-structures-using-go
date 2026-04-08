package linkedlist

func reverseLinkedList(head *Node) *Node {
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

func getKthNode(temp *Node, k int) *Node {
	k -= 1
	for temp != nil && k > 0 {
		k -= 1
		temp = temp.Next
	}
	return temp
}

func ReverseLinkedListByKGroup(head *Node, k int) *Node {
	// problem statement
	// given a head of the linked list and a positive integer known as k
	// you have to rotate the linked list individually in the group   and leave the remaining   nodes
	// if the k is not multiple of length of the linked list
	//
	//1->2->3->4->5
	curr := head
	var prev *Node = nil
	for curr != nil {
		kthNode := getKthNode(curr, k)
		if kthNode == nil {
			if prev != nil {
				prev.Next = curr
			}
			break

		}
		nextNode := kthNode.Next
		kthNode.Next = nil
		newHeading := reverseLinkedList(curr)
		if curr == head {
			head = newHeading
		} else {
			prev.Next = newHeading
		}

		prev = curr
		curr = nextNode

	}
	return head

}
