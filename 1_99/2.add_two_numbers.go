package __99

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	num, current := 0, head
	for l1 != nil || l2 != nil || num != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = new(ListNode)
		current.Next.Val = (n1 + n2 + num) % 10
		current = current.Next
		num = (n1 + n2 + num) / 10
	}
	return head.Next
}
