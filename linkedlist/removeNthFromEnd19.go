package linkedlist

import "leetcode-go/model"

/**
19. 删除链表的倒数第N个节点
https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。
*/

// 使用2次遍历
func removeNthFromEnd(head *model.ListNode, n int) *model.ListNode {
	if head == nil {
		return head
	}

	// 1 获取链表长度
	var ptr = head
	var len = 1
	for ptr.Next != nil {
		len += 1
		ptr = ptr.Next
	}

	// 2 获取倒数第N个节点的前一个节点node
	var node *model.ListNode = nil
	ptr = head
	for i := 0; i < len-n; i++ {
		node = ptr
		ptr = ptr.Next
	}

	// 3 通过前一个节点node，删除第N个节点
	if node == nil {
		return head.Next
	}
	var temp = node.Next
	node.Next = temp.Next
	temp.Next = nil

	return head
}

// 使用1次遍历
// 1->2->3->4->5, 和 n = 2.
func removeNthFromEnd2(head *model.ListNode, n int) *model.ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return nil
	}

	// 1 先让temp指针走n步
	var temp = head
	for i := 0; i < n; i++ {
		temp = temp.Next
	}

	// 如果temp为空，说明要删除第1个节点
	if temp == nil {
		return head.Next
	}

	// 2 再让temp一直走到结尾，同时ptr从head处跟随着，获得倒数第N个节点之前的节点
	var ptr = head
	for temp.Next != nil {
		ptr = ptr.Next
		temp = temp.Next
	}

	// 3 通过前一个节点node，删除第N个节点
	temp = ptr.Next
	ptr.Next = temp.Next
	temp.Next = nil

	return head
}
