package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetList(list []int) *ListNode {
	head := &ListNode{list[0], nil}
	curr := head
	for i := 1; i < len(list); i++ {
		curr.Next = &ListNode{list[i], nil}
		curr = curr.Next
	}
	return head
}
