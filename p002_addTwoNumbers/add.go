package add

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head, one, two := &ListNode{}, l1, l2
	tail := head
	carry := 0
	for one != nil && two != nil {
		sum := 0
		if one != nil {
			sum += one.Val
			one = one.Next
		}
		if two != nil {
			sum += two.Val
			two = two.Next
		}
		sum += carry
		carry = 0
		if sum >= 10 {
			sum %= 10
			carry = 1
		}
		tail.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
		tail = tail.Next
	}

	if carry > 0 {
		tail.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return head.Next
}
