package linkedlist

import (
	"fmt"
	"leetcode-go/model"
	"testing"
)

func Test(t *testing.T) {
	var head = &model.ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next = &model.ListNode{
		Val:  2,
		Next: nil,
	}
	head.Next.Next = &model.ListNode{
		Val:  3,
		Next: nil,
	}
	var newHead = removeNthFromEnd(head, 3)
	var str = model.StringifyLinkedList(newHead)
	fmt.Println("结果：" + str)
}

func Test2(t *testing.T) {
	var head = &model.ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next = &model.ListNode{
		Val:  2,
		Next: nil,
	}
	//head.Next.Next = &model.ListNode{
	//	Val:  3,
	//	Next: nil,
	//}
	var newHead = removeNthFromEnd2(head, 1)
	var str = model.StringifyLinkedList(newHead)
	fmt.Println("结果：" + str)
}
