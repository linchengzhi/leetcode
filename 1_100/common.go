package __100

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func FmtDp(data [][]int) {
	for _, val := range data {
		fmt.Println(val)
	}
}
