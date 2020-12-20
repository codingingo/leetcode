package main

// https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	// 解题思路：将结点的Child看作左子树根结点，Next看作右子树根结点
	// 最终扁平化链表即是“先序遍历”
	return flattenByIteration(root)
	// return flattenByRecursion(root)
}

func flattenByIteration(root *Node) *Node {
	if root == nil || (root.Child == nil && root.Next == nil) {
		return root
	}
	// head表示扁平化后的链表
	head := &Node{}
	// 扁平化后链表head的尾结点
	tail := head

	// 栈用于完成遍历，curr循环变量表示每次处理的结点
	stack, curr := []*Node{root}, (*Node)(nil)

	for len(stack) != 0 {
		// 获取栈中最后压入的结点
		curr = stack[len(stack)-1]
		// 将最后压入的结点弹出
		stack = stack[:len(stack)-1]

		// 修正双向链表的指针
		tail.Next = curr
		curr.Prev = tail
		tail = curr

		// 将curr当前结点的“右子树根结点”压入栈中
		if curr.Next != nil {
			stack = append(stack, curr.Next)
		}
		// 将curr当前结点的“左子树根结点”压入栈中
		if curr.Child != nil {
			stack = append(stack, curr.Child)
			// 清理标记
			curr.Child = nil
		}
	}
	// 清理真实头结点的prev指针
	head.Next.Prev = nil
	return head.Next
}

func flattenByRecursion(root *Node) *Node {
	if root == nil || (root.Child == nil && root.Next == nil) {
		return root
	}
	// 扁平化链表的虚头结点
	head := &Node{}
	// 利用递归方法添加到head
	recursion(head, root)
	// 清理真实头结点的prev指针
	head.Next.Prev = nil
	return head.Next
}

func recursion(tail, root *Node) *Node {
	// 所有结点都已添加到扁平化链表
	if root == nil {
		return tail
	}
	// 将当前结点root添加到tail，修正双向指针
	tail.Next = root
	root.Prev = tail
	tail = root
	// 当前结点root的Next指针会在“左子树根结点”的递归调用中修改，故暂存之
	next := root.Next
	// 将root的“左子树根结点”添加到tail
	tail = recursion(tail, root.Child)
	// 清理标记，因为扁平化后的链表中的结点不允许再有Child
	root.Child = nil
	// 将root的“右子树根结点”添加到tail
	return recursion(tail, next)
}
