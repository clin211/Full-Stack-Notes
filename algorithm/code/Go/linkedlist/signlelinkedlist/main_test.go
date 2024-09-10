package main

import (
	"fmt"
	"testing"
)

// Int 类型实现 Comparable 接口
type Int int

func (a Int) Compare(b Int) bool {
	return a == b
}

// Print 打印链表
func (list *LinkedList[T]) Print() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

// 测试链表的基本功能
func TestLinkedList(t *testing.T) {
	list := NewLinkedList[Int]()

	// 测试插入头部
	if !list.InsertHead(1) {
		t.Errorf("InsertHead failed")
	}
	list.Print()
	if list.length != 1 {
		t.Errorf("Expected length 1, got %d", list.length)
	}

	// 测试插入尾部
	if !list.InsertTail(2) {
		t.Errorf("InsertTail failed")
	}
	if list.length != 2 {
		t.Errorf("Expected length 2, got %d", list.length)
	}
	list.Print()

	// 测试查找
	node := list.Find(1)
	if node == nil || node.GetValue() != 1 {
		t.Errorf("Find failed, expected 1, got %v", node)
	}
	fmt.Println("node", node)
	// 测试更新
	if !list.Update(1, 3) {
		t.Errorf("Update failed")
	}
	node = list.Find(3)
	if node == nil {
		t.Errorf("Find failed for updated value, expected 3, got nil")
	}
	list.Print()

	// 测试删除
	if !list.Delete(node) {
		t.Errorf("Delete failed")
	}
	if list.Find(3) != nil {
		t.Errorf("Expected to not find 3 after deletion")
	}
	list.Print()
	// 测试链表长度
	if list.length != 1 {
		t.Errorf("Expected length 1 after deletion, got %d", list.length)
	}
	fmt.Println("list", list, "length", list.length)
	// 测试插入后删除
	list.InsertTail(4)
	list.InsertTail(5)
	list.Print()
	if list.length != 3 {
		t.Errorf("Expected length 3 after insertions, got %d", list.length)
	}

	if !list.Delete(list.Find(4)) {
		t.Errorf("Delete failed for value 4")
	}
	if list.length != 2 {
		t.Errorf("Expected length 2 after deletion, got %d", list.length)
	}
	list.Print()
}
