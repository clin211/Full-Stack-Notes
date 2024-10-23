package stack

import (
	"fmt"
	"testing"
)

// 定义一个简单的类型，实现 Comparable 接口
type Int int

func (a Int) Compare(other Int) bool {
	return a == other
}

func TestLinkedListStack(t *testing.T) {
	stack := NewLinkedListStack[Int]()

	// 测试初始状态
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
	if size := stack.Size(); size != 0 {
		t.Errorf("Expected size to be 0, got %d", size)
	}

	// 测试入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// 测试栈的状态
	if stack.IsEmpty() {
		t.Errorf("Expected stack to not be empty")
	}
	if size := stack.Size(); size != 3 {
		t.Errorf("Expected size to be 3, got %d", size)
	}

	// 测试栈顶元素
	if top, ok := stack.Top(); !ok || top != 3 {
		t.Errorf("Expected top element to be 3, got %d", top)
	}

	// 测试出栈
	if popped, ok := stack.Pop(); !ok || popped != 3 {
		t.Errorf("Expected popped element to be 3, got %d", popped)
	}
	if size := stack.Size(); size != 2 {
		t.Errorf("Expected size to be 2 after pop, got %d", size)
	}

	// 测试再次出栈
	if popped, ok := stack.Pop(); !ok || popped != 2 {
		t.Errorf("Expected popped element to be 2, got %d", popped)
	}

	// 测试栈顶元素
	if top, ok := stack.Top(); !ok || top != 1 {
		t.Errorf("Expected top element to be 1, got %d", top)
	}

	// 测试清空栈
	stack.Clear()
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after clear")
	}
	if size := stack.Size(); size != 0 {
		t.Errorf("Expected size to be 0 after clear, got %d", size)
	}

	// 测试从空栈出栈
	if popped, ok := stack.Pop(); ok || popped != 0 {
		t.Errorf("Expected popped element from empty stack to be 0 and ok to be false, got %d", popped)
	}

	// 测试栈的打印（可选）
	fmt.Println("Current stack state:")
	stack.Print()
}
