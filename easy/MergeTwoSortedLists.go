// 2021.06.27

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func appendToList(head, node *ListNode) {
	cursor := head
	for {
		if nil == cursor.Next {
			break
		}
		cursor = cursor.Next
	}
	cursor.Next = node
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 && nil == l2 {
		return nil
	}
	if nil == l1 && l2 != nil {
		return l2
	}
	if l1 != nil && nil == l2 {
		return l1
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	for {
		if nil == l1 || nil == l2 {
			break
		}
		if l1.Val < l2.Val {
			tmp := l1.Next
			l1.Next = nil
			appendToList(head, l1)
			l1 = tmp
		} else {
			tmp := l2.Next
			l2.Next = nil
			appendToList(head, l2)
			l2 = tmp
		}
	}
	if l1 != nil {
		appendToList(head, l1)
	}
	if l2 != nil {
		appendToList(head, l2)
	}

	return head.Next
}
