package recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

func IntsToList(list []int) *ListNode {
	head := new(ListNode)
	current := head
	for _, val := range list {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head.Next
}

func ListToInt(list *ListNode) []int {
	nums := []int{}
	for list != nil {
		nums = append(nums, list.Val)
		list = list.Next
	}
	return nums
}
