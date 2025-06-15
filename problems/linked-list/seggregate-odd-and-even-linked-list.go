package linkedlist

func SeggregateOddAndEvenBruteForce(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	ans := []int{}
	curr := head
	for curr != nil && curr.Next != nil {
		ans = append(ans, curr.Data)
		curr = curr.Next.Next
	}
	if curr != nil {
		ans = append(ans, curr.Data)
	}
	curr = head.Next
	for curr != nil && curr.Next != nil {
		ans = append(ans, curr.Data)
		curr = curr.Next.Next
	}

	index := 0
	curr = head
	for curr != nil {
		curr.Data = ans[index]
		index++
		curr = curr.Next
	}
	return head
}
func SeggregateOddAndEvenOptimalAproach(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := head.Next
	for even != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head

}
