package main

import (
	"errors"
	"fmt"
)

type CircularQueue struct {
	data  []int // 存储数据
	head  int   // 头指针
	tail  int   // 尾指针
	count int   // 当前元素个数
	size  int   // 队列容量
}

// 创建一个新的循环队列
func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		data:  make([]int, size),
		head:  0,
		tail:  0,
		count: 0,
		size:  size,
	}
}

// 入队操作
func (q *CircularQueue) Enqueue(value int) error {
	if q.count == q.size {
		return errors.New("队列已满")
	}
	q.data[q.tail] = value
	q.tail = (q.tail + 1) % q.size
	q.count++
	return nil
}

// 打印队列内容
func (q *CircularQueue) Print() {
	for i := 0; i < q.count; i++ {
		index := (q.head + i) % q.size
		fmt.Print(q.data[index], " ")
	}
	fmt.Println()
}

func main() {
	queue := NewCircularQueue(5)

	// 入队操作示例
	for i := 1; i <= 5; i++ {
		if err := queue.Enqueue(i); err != nil {
			fmt.Println(err)
		}
	}

	queue.Print() // 输出: 1 2 3 4 5

	// 尝试再入队一个元素
	if err := queue.Enqueue(6); err != nil {
		fmt.Println(err) // 输出: 队列已满
	}
}
