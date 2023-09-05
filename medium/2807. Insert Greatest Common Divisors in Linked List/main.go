package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil {
		gcdNode := &ListNode{
			Val:  gcd(slow.Val, fast.Val),
			Next: fast,
		}

		slow.Next = gcdNode

		slow = fast
		fast = slow.Next
	}

	return head
}
