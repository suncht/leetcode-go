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
