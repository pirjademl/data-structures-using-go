package linkedlist

func ArrToLinkedList(arr []int) *Node {
	head := &Node{arr[0], nil}
	var curr *Node
	curr = head
	for i := 1; i < len(arr); i++ {
		temp := &Node{arr[i], nil}
		curr.Next = temp
		curr = curr.Next
	}
	return head
}
