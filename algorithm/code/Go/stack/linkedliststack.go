package stack

import "fmt"

// Comparable 接口用于比较泛型类型
type Comparable[T any] interface {
	Compare(other T) bool
}

// 定义链表节点，T 为类型参数
type Node[T any] struct {
	data T        // 数据域
	next *Node[T] // 指向下一个节点
}

// 定义链表结构，T 为类型参数
type LinkedListStack[T Comparable[T]] struct {
	head   *Node[T] // 头节点
	length uint     // 链表长度
}

// 用于获取类型 T 的零值
func zeroValue[T any]() T {
	var zero T
	return zero
}

// 创建新节点
func NewListNode[T Comparable[T]](v T) *Node[T] {
	return &Node[T]{
		data: v,
		next: nil,
	}
}

// 创建新链表
func NewLinkedListStack[T Comparable[T]]() *LinkedListStack[T] {
	return &LinkedListStack[T]{
		head:   NewListNode[T](zeroValue[T]()),
		length: 0,
	}
}

// 判断链表是否为空
func (list *LinkedListStack[T]) IsEmpty() bool {
	return list.length == 0
}

// 链表入栈
func (list *LinkedListStack[T]) Push(v T) {
	list.head = &Node[T]{
		data: v,
		next: list.head,
	}
	list.length++
}

// 链表出栈
func (list *LinkedListStack[T]) Pop() (data T, ok bool) {
	if list.IsEmpty() {
		return zeroValue[T](), false
	}
	current := list.head
	// 获取链表头部的下一个节点
	list.head = current.next
	list.length--
	return current.data, true
}

// 链表栈顶
func (list *LinkedListStack[T]) Top() (data T, ok bool) {
	if list.IsEmpty() {
		return zeroValue[T](), false
	}
	return list.head.data, true
}

// 链表长度
func (list *LinkedListStack[T]) Size() uint {
	return uint(list.length)
}

// 清空
func (list *LinkedListStack[T]) Clear() {
	list.head = NewListNode[T](zeroValue[T]())
	list.length = 0
}

// 打印
func (list *LinkedListStack[T]) Print() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}
