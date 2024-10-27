package queue

import (
	"testing"
)

func TestArrayQueue_Enqueue(t *testing.T) {
	queue := NewArrayQueue[int](3)

	// 队列应为空
	if !queue.IsEmpty() {
		t.Errorf("队列初始化后应为空")
	}

	// 入队元素
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// 队列应为满
	if !queue.IsFull() {
		t.Errorf("队列应已满")
	}

	// 尝试再入队，队列满则无法入队
	queue.Enqueue(4)
	if queue.Size() != 3 {
		t.Errorf("队列已满时不能继续入队")
	}
}

func TestArrayQueue_Dequeue(t *testing.T) {
	queue := NewArrayQueue[int](32)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// 出队测试
	item, ok := queue.Dequeue()
	if !ok || item != 1 {
		t.Errorf("期望出队元素为 1，得到 %v", item)
	}

	item, ok = queue.Dequeue()
	if !ok || item != 2 {
		t.Errorf("期望出队元素为 2，得到 %v", item)
	}

	// 再次出队
	item, ok = queue.Dequeue()
	if !ok || item != 3 {
		t.Errorf("期望出队元素为 3，得到 %v", item)
	}

	// 出队空队列
	_, ok = queue.Dequeue()
	if ok {
		t.Errorf("空队列出队应返回 false")
	}
}

func TestArrayQueue_Front(t *testing.T) {
	queue := NewArrayQueue[int](32)

	// 队列为空时，Front() 应返回零值
	if queue.Front() != 0 {
		t.Errorf("空队列 Front 应返回元素类型的零值")
	}

	queue.Enqueue(1)
	queue.Enqueue(2)

	// Front 应返回队头元素，不移除
	if queue.Front() != 1 {
		t.Errorf("期望 Front 返回 1，得到 %v", queue.Front())
	}

	// 出队后，Front 应返回下一个元素
	queue.Dequeue()
	if queue.Front() != 2 {
		t.Errorf("期望 Front 返回 2，得到 %v", queue.Front())
	}
}

func TestArrayQueue_IsEmptyAndIsFull(t *testing.T) {
	queue := NewArrayQueue[int](3)

	// 初始队列应为空
	if !queue.IsEmpty() {
		t.Errorf("初始化队列应为空")
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// 队列应为满
	if !queue.IsFull() {
		t.Errorf("队列应已满")
	}

	// 出队后应非满
	queue.Dequeue()
	if queue.IsFull() {
		t.Errorf("出队后队列不应为满")
	}

	// 再次入队，队列应满
	queue.Enqueue(4)
	if !queue.IsFull() {
		t.Errorf("入队后队列应为满")
	}
}

func TestArrayQueue_Size(t *testing.T) {
	queue := NewArrayQueue[int](32)

	// 初始大小应为 0
	if queue.Size() != 0 {
		t.Errorf("初始队列大小应为 0")
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// 入队三个元素，大小应为 3
	if queue.Size() != 3 {
		t.Errorf("入队后队列大小应为 3，得到 %v", queue.Size())
	}

	queue.Dequeue()

	// 出队一个元素，大小应为 2
	if queue.Size() != 2 {
		t.Errorf("出队后队列大小应为 2，得到 %v", queue.Size())
	}
}
