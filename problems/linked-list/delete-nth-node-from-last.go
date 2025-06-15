package linkedlist

func DeleteNthNodeFromLast(head *Node, n int) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// one edge case that if you have asked to delete the head
	if fast == nil {
		return head.Next
	}
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	//slow will be at one step before the node to be deletedo
	slow.Next = slow.Next.Next
	return head
}
