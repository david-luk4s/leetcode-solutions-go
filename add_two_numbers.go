/*
 * https://leetcode.com/problems/add-two-numbers/
 *
 * 2. Add Two Numbers
 * Medium
 *
 * Runtime: 7 ms, faster than 93.04% of Go online submissions for Add Two Numbers.
 * Memory Usage: 4.5 MB, less than 81.54% of Go online submissions for Add Two Numbers.
 *
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order, and each of their nodes contains a single digit.
 * Add the two numbers and return the sum as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 *
 * Example 1:
 *
 * Input: l1 = [2,4,3], l2 = [5,6,4]
 * Output: [7,0,8]
 * Explanation: 342 + 465 = 807.
 *
 */
/**
 Definition for singly-linked list.
 type ListNode struct {
	Val  int
	Next *ListNode
}
**/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	curr := head
	n1, n2, carr := 0, 0, 0

	for l1 != nil || l2 != nil || carr != 0 {
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}

		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}

		sum := n1 + n2 + carr
		carr = sum / 10
		num := sum % 10
		n1, n2 = 0, 0

		curr.Next = &ListNode{Val: num, Next: nil}
		curr = curr.Next
	}

	return head.Next
}
