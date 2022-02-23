package main

import "fmt"

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	//链表长度
	n := 1

	cur := head

	//计算链表的长度
	for cur.Next != nil {
		n++
		cur = cur.Next
	}
	// fmt.Printf("链表长度 = %d\n", n)

	//计算有效移动的步数
	step := (k%n) % n
	// fmt.Printf("有效移动步数 = %d\n", step)

	//如果步数为0, 直接返回原链表
	// if step == 0 {
	// 	return head
	// }

	//构造成环形链表, 最后结点的Next指向head
	cur.Next = head

	//cur再指向head
	cur = head

	//查找移动后对应的尾结点
	for i := 0; i < n - step - 1; i++ {
		// fmt.Println("移动步数", i)
		cur = cur.Next
	}

	fmt.Println("尾结点 = ", cur.Val)

	//确定新的头结点, 此时cur是尾结点, 它的Next就是新的头结点
	newHead := cur.Next
	//断开尾结点指向头结点的链接
	cur.Next = nil

	return newHead
}



// @lc code=end
