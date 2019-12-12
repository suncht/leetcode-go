package linkedlist

import "leetcode-go/model"

/**
23. 合并K个排序链表
https://leetcode-cn.com/problems/merge-k-sorted-lists/
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

func mergeKLists(lists []*model.ListNode) *model.ListNode {
	if lists == nil {
		return nil
	}
	var length = len(lists)
	if length == 0 {
		return nil
	}

	// 整体时间复杂度  O(KlogK)
	// 两两组合合并， O(logN)
	for length > 1 {
		for i := 0; i < length/2; i++ {
			// 两个链表合并 O(K)
			lists[i] = mergeTwoList(lists[i], lists[length-1-i])
		}
		length = (length + 1) / 2
	}

	return lists[0]
}

/**
合并2个排序链表
*/
func mergeTwoList(list1 *model.ListNode, list2 *model.ListNode) *model.ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	if list1 == nil && list2 != nil {
		return list2
	}

	var newList = &model.ListNode{} // 使用空节点
	var ptr0 = newList
	var ptr1 = list1
	var ptr2 = list2
	// 1 顺序比较2个链接的节点，用ptr0将节点串联起来
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val <= ptr2.Val {
			ptr0.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			ptr0.Next = ptr2
			ptr2 = ptr2.Next
		}
		ptr0 = ptr0.Next
	}

	// 2 如果第一个链表还有数据， 而第二个链表没有数据，则直接挂接第一个链表后续的节点
	if ptr1 != nil && ptr2 == nil {
		ptr0.Next = ptr1
	}

	// 3 如果第一个链表还有数据， 而第二个链表没有数据，则直接挂接第二个链表后续的节点
	if ptr1 == nil && ptr2 != nil {
		ptr0.Next = ptr2
	}

	// 4 去掉第一个空节点
	emptyNode := newList
	newList = newList.Next
	emptyNode.Next = nil

	return newList
}
