package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
}

// 在循环链表的末尾插入新节点
func (cll *CircularLinkedList) Insert(data int) {
	newNode := &Node{data: data}
	if cll.head == nil {
		cll.head = newNode
		newNode.next = cll.head
	} else {
		current := cll.head
		for current.next != cll.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = cll.head
	}
}

// 根据传入的data，删除循环链表中的节点
// - **链表为空**：为空则直接退出。
// - **链表非空**：遍历找出节点
//   - 如果是头节点：
//     1.只有一个节点，直接将 head 设置为 `NULL。`
//     2.多个节点，找到最后一个节点，并将其更新前一个节点的 next 指针，跳过要删除的节点。
//   - 非头节点
//     1.直接将要删除节点的前一个节点的 next 指针指向要删除节点的下一个节点（也就是跳过要删除的节点）
func (cll *CircularLinkedList) Delete(data int) error {
	if cll.head == nil {
		return errors.New("list is empty")
	}

	// 如果是头节点
	if cll.head.data == data {
		if cll.head.next == cll.head {
			// 只有一个节点
			cll.head = nil
		} else {
			// 多个节点
			current := cll.head
			for current.next != cll.head {
				current = current.next
			}
			current.next = cll.head.next
			cll.head = cll.head.next
		}
		return nil
	}

	// 非头节点
	current := cll.head
	for current.next != cll.head {
		if current.next.data == data {
			current.next = current.next.next
			return nil
		}
		current = current.next
	}

	return errors.New("data not found")
}

// 打印循环链表中的所有节点
func (cll *CircularLinkedList) Print() {
	if cll.head == nil {
		fmt.Println("List is empty")
		return
	}
	current := cll.head
	for {
		fmt.Print(current.data, " -> ")
		current = current.next
		if current == cll.head {
			break
		}
	}
	fmt.Println("(back to head)")
}

func main() {
	cll := &CircularLinkedList{}

	cll.Insert(1)
	cll.Insert(2)
	cll.Insert(3)

	cll.Print() // 输出: 1 -> 2 -> 3 -> (back to head)
}
