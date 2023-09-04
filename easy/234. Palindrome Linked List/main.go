package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	head = prev

	return head
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	revList := reverseList(slow)

	for pointer := revList; pointer != nil; pointer, head = pointer.Next, head.Next {
		if pointer.Val != head.Val {
			return false
		}
	}

	return true
}
