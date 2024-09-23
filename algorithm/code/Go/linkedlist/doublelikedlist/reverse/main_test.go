package main

import (
	"testing"
)

// TestReverse 测试反转双向链表的功能
func TestReverse(t *testing.T) {
	list := &List{}

	test_data := []int{1, 5, 8, 9, 3, 6}
	expected := []int{6, 3, 9, 8, 5, 1}

	// 添加元素
	for _, v := range test_data {
		list.append(v)
	}

	// 反转链表
	list.reverse()

	current := list.head
	for i, v := range expected {
		if current == nil || current.data != v {
			t.Errorf("Expected %d at position %d, got %d", v, i, current.data)
		}
		current = current.next
	}

	// 确保链表末尾为 nil
	if current != nil {
		t.Error("Expected end of list to be nil, but it was not")
	}
}

// append 方法用于在链表末尾添加节点
func (list *List) append(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	list.length++
}
