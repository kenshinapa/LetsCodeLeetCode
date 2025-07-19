// Package addtwonumbers - Problem 2. Add Two Numbers
package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummyHead.Next
}
