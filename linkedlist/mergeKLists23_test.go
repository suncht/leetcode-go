package linkedlist

import (
	"fmt"
	"leetcode-go/model"
	"testing"
)

func TestMergeKLists23(t *testing.T) {
	var list1 = model.GenerateLinkedList(4, 2)
	var list2 = model.GenerateLinkedList(3, 1)

	fmt.Println(model.StringifyLinkedList(list1))
	fmt.Println(model.StringifyLinkedList(list2))

	var list = mergeTwoList(list1, list2)
	fmt.Println(model.StringifyLinkedList(list))
}

func TestMergeKLists23_res(t *testing.T) {
	var list1 = model.GenerateLinkedList(4, 2)
	var list2 = model.GenerateLinkedList(3, 3)
	var list3 = model.GenerateLinkedList(2, 5)

	fmt.Println(model.StringifyLinkedList(list1))
	fmt.Println(model.StringifyLinkedList(list2))
	fmt.Println(model.StringifyLinkedList(list3))

	var lists = []*model.ListNode{list1, list2, list3}

	var list = mergeKLists(lists)
	fmt.Println(model.StringifyLinkedList(list))
}
