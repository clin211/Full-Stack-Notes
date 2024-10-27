package queue

type ArrayQueue[T any] struct {
	data     []T
	capacity int // 队列容量
	front    int // 队头指针
	rear     int // 队尾指针
	size     int // 队列长度
}

type ArrayQueueMethod[T any] interface {
	Enqueue(T)
	Dequeue() (T, bool)
	Front() T
	IsEmpty() bool
	IsFull() bool
	Size() int
}

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		data:     make([]T, capacity),
		capacity: capacity,
		front:    0,
		rear:     -1,
		size:     0,
	}
}

var _ ArrayQueueMethod[int] = (*ArrayQueue[int])(nil)

// 入队
func (aq *ArrayQueue[T]) Enqueue(data T) {
	// 判断队列是否已满
	if aq.IsFull() {
		return
	}
	aq.rear = (aq.rear + 1) % aq.capacity
	aq.data[aq.rear] = data
	aq.size++
}

// 出队
func (aq *ArrayQueue[T]) Dequeue() (T, bool) {
	// 判断队列是否为空
	if aq.IsEmpty() {
		var zero T
		return zero, false
	}
	// 队头元素出队
	item := aq.data[aq.front]
	// 更新对头
	aq.front = (aq.front + 1) % aq.capacity
	aq.size--
	return item, true
}

// 队头
func (aq *ArrayQueue[T]) Front() T {
	if aq.IsEmpty() {
		var zero T
		return zero
	}
	return aq.data[aq.front]
}

// 判断队列是否为空
func (aq *ArrayQueue[T]) IsEmpty() bool {
	return aq.size == 0
}

// 判断队列是否已满
func (aq *ArrayQueue[T]) IsFull() bool {
	return aq.size == aq.capacity
}

// 队列长度
func (aq *ArrayQueue[T]) Size() int {
	return aq.size
}
