package deletenode

import linkedlist "github.com/dsa/problems/linked-list"

func DeleteTail(head *linkedlist.Node) *linkedlist.Node {
	// I have to stop one node before the DeleteTail
	if head == nil || head.Next == nil {
		return nil
	}
	curr := head

	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
	return head

}
