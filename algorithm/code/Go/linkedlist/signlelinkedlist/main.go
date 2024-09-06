package main

import (
	"fmt"
)

// 定义链表节点结构
type Node struct {
	Data int
	Next *Node
}

// 向链表中插入元素
func Insert(head *Node, data int, position int) *Node {
	if position < 0 {
		return head
	}
	// 创建临时节点 temp
	temp := head

	// 找到要插入位置的前一个节点
	for i := 1; i < position; i++ {
		if temp == nil {
			fmt.Println("插入位置无效")
			return head
		}
		temp = temp.Next
	}

	// 创建插入节点 c
	newNode := &Node{Data: data}

	// 向链表中插入节点
	if temp != nil {
		newNode.Next = temp.Next
		temp.Next = newNode
	} else {
		// 如果 temp 是 nil，说明要插入在链表末尾
		newNode.Next = nil
		head = newNode
	}

	return head
}

// 打印链表
func PrintList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// 创建一个单链表
	head := &Node{Data: 1}
	head.Next = &Node{Data: 2}
	head.Next.Next = &Node{Data: 3}

	fmt.Println("原链表:")
	PrintList(head)

	// 在链表的第2个位置插入元素 4
	head = Insert(head, 4, 2)
	fmt.Println("插入元素 4 后的链表:")
	PrintList(head)

	// 在链表的末尾插入元素 5
	head = Insert(head, 5, 5)
	fmt.Println("插入元素 5 后的链表:")
	PrintList(head)
}
