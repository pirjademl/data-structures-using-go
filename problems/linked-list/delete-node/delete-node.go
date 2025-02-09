package deletenode

import linkedlist "github.com/dsa/problems/linked-list"

func DeleteNode(node *linkedlist.Node) {
	node.Data = node.Next.Data
	node.Next = node.Next.Next
}
