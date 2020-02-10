package questions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 != nil {
			return l1
		} else if l2 != nil {
			return l2
		} else {
			return nil
		}
	}

	mergedList := &ListNode{
		Val:  0,
		Next: nil,
	}
	p1 := l1
	p2 := l2
	for p1.Next != nil || p2.Next != nil {
		if l1.Val > l2.Val {
			mergedList.Next = p2
			p2 = p2.Next
		} else {
			mergedList.Next = p1
			p1 = p1.Next
		}
		mergedList = mergedList.Next
	}
	mergedList.Next = map[bool]*ListNode{true: p1, false: p2}[p1 != nil]

	return mergedList.Next
}
