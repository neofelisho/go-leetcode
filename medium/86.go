package medium

func partition(head *ListNode, x int) *ListNode {
	// ex. list->1->4->3->2->5->2, x=3
	// s1: list->1->4->3->2->5->2
	// s2: list->1->3->2->5->2
	//     temp->4
	// s3: list->1->2->5->2
	//     temp->4->3
	// ...
	// fi: list->1->2->2
	//     temp->4->3->5

	if head == nil || head.Next == nil {
		return head
	}
	current := head
	var (
		temp     *ListNode
		tempTail *ListNode
		prev     *ListNode
	)
	for current != nil {
		next := current.Next
		if current.Val < x {
			prev = current
		} else {
			if temp == nil && tempTail == nil {
				temp = current
				tempTail = current
			} else {
				tempTail.Next = current
				tempTail = tempTail.Next
			}
			tempTail.Next = nil

			if prev == nil {
				head = next
			} else {
				prev.Next = next
			}
		}
		current = next
	}

	if prev == nil {
		head = temp
	} else {
		prev.Next = temp
	}
	return head
}
