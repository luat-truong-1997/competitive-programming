/**
* Problem: https://leetcode.com/problems/reverse-nodes-in-even-length-groups/submissions/870771363/
* Reversed as what they said.
 */
package main

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

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	count := 0
	group := 1
	node := head
	var prev *ListNode = nil
	var startGroup *ListNode = head
	var prevStartGroup *ListNode = nil
	for node != nil {
		count++
		prev = node
		node = node.Next
		if count == group {
			if count%2 == 0 {
				// Visit next group
				root, lastNode, nextOfLast := reverse(startGroup, 0, count)
				lastNode.Next = nextOfLast
				prevStartGroup.Next = root.Next
				// Currently node is last node
				prevStartGroup = lastNode
				startGroup = lastNode.Next
			} else {
				startGroup = node
				prevStartGroup = prev
			}
			count = 0
			group++
		}
	}
	if (count != 0) && (count%2 == 0) {
		root, lastNode, nextOfLast := reverse(startGroup, 0, count)
		lastNode.Next = nextOfLast
		prevStartGroup.Next = root.Next
	}
	return head
}

func reverse(node *ListNode, index int, group int) (*ListNode, *ListNode, *ListNode) {
	if index == group-1 {
		return &ListNode{Next: node}, node, node.Next
	}
	root, prevNode, nextOfLast := reverse(node.Next, index+1, group)
	prevNode.Next = node
	return root, node, nextOfLast
}
