package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	sum := 0
	result := &ListNode{
		Val:  0,
		Next: nil,
	}
	pointer := result

	for left := head.Next; left != nil; left = left.Next {
		if left.Val == 0 {
			pointer.Next = &ListNode{
				Val:  sum,
				Next: nil,
			}
			pointer = pointer.Next
			sum = 0
		} else {
			sum += left.Val
		}
	}

	return result.Next
}
