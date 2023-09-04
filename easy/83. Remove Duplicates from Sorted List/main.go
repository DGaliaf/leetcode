package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	res := head

	if head == nil {
		return head
	}

	for head != nil {
		for head.Next != nil && head.Next.Val == head.Val {
			head.Next = head.Next.Next
		}

		head = head.Next
	}

	return res
}
