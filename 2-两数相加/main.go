package main

import bean "LeetCode"

type ListNode = bean.ListNode

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/**
	思路:	利用递归
	*/
	return add(l1, l2, 0)
}

func add(l1 *ListNode, l2 *ListNode, z int) *ListNode {

	// 如果都为空代表结束,返回nil
	if l1 == nil && l2 == nil && z == 0 {
		return nil
	}
	// 如果l1为空 则新建一个
	if l1 == nil {
		l1 = new(ListNode)
	}
	// 如果l2为空 则新建一个
	if l2 == nil {
		l2 = new(ListNode)
	}
	// 计算总和 z为进位的值
	v := l1.Val + l2.Val + z
	result := new(ListNode)
	// 值等于和对10取余
	result.Val = v % 10
	// 下一个值进行递归 z为何对10整除
	result.Next = add(l1.Next, l2.Next, v/10)
	return result
}
