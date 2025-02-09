package deletenode

import linkedlist "github.com/dsa/problems/linked-list"

func DeleteHead(head *linkedlist.Node) *linkedlist.Node {
	head = head.Next
	return head
}
