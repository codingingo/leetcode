package main

// https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
// ListNode defines for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	return ByRecursion(head, k)
}

func ByRecursion(head *ListNode, k int) int {
	_, val := Recursion(head, k)
	return val
}

func Recursion(head *ListNode, k int) (int, int) {
	// 递归基
	if head == nil {
		return -1, -1
	}
	// 返回值表示，（倒数计数，对应链表结点的值）
	prevK, val := Recursion(head.Next, k)
	if prevK == -1 {
		// 表示当前函数栈帧是倒数第一个
		return 1, head.Val
	} else if prevK == k {
		// 前一个栈帧已经是第K个，透穿前一次结果即可
		return prevK, val
	} else {
		// 在前一栈帧返回值上调整倒数计数值
		// 并将结点值替换为当前结点的值
		return prevK + 1, head.Val
	}
}

func BySliceIteration(head *ListNode, k int) int {
	// 使用切片来存储链表值
	vals := []int{}
	for node := head; node != nil; node = node.Next {
		vals = append(vals, node.Val)
	}
	// 利用切片/数组的随机访问属性获取倒数第K个
	return vals[len(vals)-k]
}

func ByOnePointerIteration(head *ListNode, k int) int {
	// 获取链表长度
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	// 向前走n-k步

	for i := 0; i < n-k; i++ {
		head = head.Next
	}
	return head.Val
}

func ByTwoPointersIteration(head *ListNode, k int) int {
	// 双指针-快慢指针法，注意起始位置
	slow, fast := head, head
	// 快指针先走k步
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	// 再快、慢指针一起走
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	// 慢指针指向的值即是倒数第K个
	return slow.Val
}
