package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetListHead(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := &ListNode{list[0], nil}
	curr := head
	for i := 1; i < len(list); i++ {
		curr.Next = &ListNode{list[i], nil}
		curr = curr.Next
	}
	return head
}
