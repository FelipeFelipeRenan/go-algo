package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {

			break
		}
	}
	slow = head
	for {
		if fast == slow {
			return slow
		}
		fast = fast.Next
		slow = slow.Next

	}
}
