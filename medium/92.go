package medium

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// ex. 1->2->3->4->5->6->7, m=3, n=5.
	//
	//         n     m
	//         v     v
	//         5->4->3
	// 1 -> 2           6->7
	// ^    ^           ^
	// list tail        current

	if m == n {
		return head
	}

	counter := 1
	current := head
	var (
		tail     *ListNode
		pointerM *ListNode
		pointerN *ListNode
		next     *ListNode
	)
	for counter <= n && current != nil {
		next = current.Next
		if counter < m {
			tail = current
		}
		if counter == m {
			pointerM = current
			pointerN = current
		}
		if counter > m && counter <= n {
			current.Next = pointerN
			pointerN = current
		}
		current = next
		counter++
	}
	if tail != nil {
		tail.Next = pointerN
	} else {
		head = pointerN
	}
	if pointerM != nil {
		pointerM.Next = current
	}
	return head
}
