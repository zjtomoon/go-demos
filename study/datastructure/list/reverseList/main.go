package main

import (
	"fmt"
)

/**
开发一个程序，此程序创建一个链表，链表上一共有10个节点，每个节点是对应的数字编号（1到10），通过编程实现从反序打印链表上节点编号的程序，
也即打出的结果是10到1。
1、需在程序中生成按编号递增的单链表，如何反序打印可灵活处理，包括不限于使用数组等。
2、注意程序执行完毕后应当释放内存。
*/

type ListNode struct {
	No   int       //编号
	Next *ListNode // 指向下一个结点
}

func main() {
	fmt.Println("生成顺序链表")
	head := &ListNode{}
	GenerateList(head, 10)
	ShowListNode(head)
	fmt.Println()
	fmt.Println("链表反转")
	reverselist := ReverseList(head)
	ShowListNode(reverselist)
}

// 实现链表翻转
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 创建顺序的链表 根据输入的结点数量生成指定的链表
func GenerateList(head *ListNode, n int) {
	for i := 1; i <= n; i++ {
		InsertNode(head, &ListNode{No: i})
	}
}

// 按照No顺序地插入结点
func InsertNode(head *ListNode, newNode *ListNode) {
	temp := head
	flag := true
	for {
		if temp.Next == nil {
			break
		} else if temp.Next.No >= newNode.No {
			break
		} else if temp.Next.No == newNode.No {
			flag = false
			break
		}
		temp = temp.Next
	}
	if !flag {
		fmt.Println("already have a no = ", newNode.No)
		return
	} else {
		newNode.Next = temp.Next
		temp.Next = newNode
	}
}

// 显示链表所有的结点
func ShowListNode(head *ListNode) {
	temp := head
	if temp.Next == nil {
		fmt.Println("nil list")
		return
	}
	for {
		fmt.Printf("[%d] ==>", temp.Next.No)
		temp = temp.Next
		if temp.Next == nil {
			break
		}
	}
}
