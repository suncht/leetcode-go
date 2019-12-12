package model

import (
	"bytes"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
获取Linked-list链表的字符串
*/
func StringifyLinkedList(head *ListNode) string {
	if head == nil {
		return ""
	}
	var buffer bytes.Buffer
	for head != nil {
		buffer.WriteString(strconv.Itoa(head.Val))

		if head = head.Next; head != nil {
			buffer.WriteString(" --> ")
		}
	}
	return buffer.String()
}

func GetLinkedListWithOne() *ListNode {
	var head = &ListNode{
		Val:  1,
		Next: nil,
	}

	return head
}
func GetLinkedListWith2() *ListNode {
	return GenerateLinkedList(2, 1)
}

func GetLinkedListWith3() *ListNode {
	return GenerateLinkedList(3, 1)
}

func GetLinkedListWith4() *ListNode {
	return GenerateLinkedList(4, 1)
}

func GenerateLinkedList(n int, incrment int) *ListNode {
	if incrment == 0 {
		incrment = 1
	}
	var head = &ListNode{
		Val:  1 * incrment,
		Next: nil,
	}
	var ptr = head
	for i := 2; i <= n; i++ {
		ptr.Next = &ListNode{
			Val:  i * incrment,
			Next: nil,
		}
		ptr = ptr.Next
	}
	return head
}
