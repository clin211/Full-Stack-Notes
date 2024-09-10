package main

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
type LinkedList[T Comparable[T]] struct {
	head   *Node[T] // 头节点
	length uint     // 链表长度
}

// 创建新节点
func NewListNode[T Comparable[T]](v T) *Node[T] {
	return &Node[T]{
		data: v,
		next: nil,
	}
}

// 创建新链表
func NewLinkedList[T Comparable[T]]() *LinkedList[T] {
	return &LinkedList[T]{
		head:   NewListNode[T](zeroValue[T]()),
		length: 0,
	}
}

// 获取下一个节点
func (list *Node[T]) GetNext() *Node[T] {
	return list.next
}

// 获取节点数据
func (list *Node[T]) GetValue() T {
	return list.data
}

// 用于获取类型 T 的零值
func zeroValue[T any]() T {
	var zero T
	return zero
}

// InsertAfter 在指定节点 p 后插入值为 v 的新节点
func (list *LinkedList[T]) InsertAfter(p *Node[T], v T) bool {
	// 如果 p 为空，则返回 false
	if p == nil {
		return false
	}

	// 创建新节点
	newNode := NewListNode(v)
	// 记录原来的下一个节点
	oldNext := p.next
	// 将新节点插入到 p 之后
	p.next = newNode
	newNode.next = oldNext

	// 链表长度加1
	list.length++
	return true
}

// 在链表头部插入节点
func (list *LinkedList[T]) InsertHead(v T) bool {
	return list.InsertAfter(list.head, v)
}

// 在链表尾部插入节点
func (list *LinkedList[T]) InsertTail(v T) bool {
	current := list.head

	for current.next != nil {
		current = current.next
	}
	return list.InsertAfter(current, v)
}

// 删除传入的节点
func (list *LinkedList[T]) Delete(p *Node[T]) bool {
	// 如果 p 为空，则返回 false
	if p == nil {
		return false
	}

	current := list.head.next
	previous := list.head

	for current != nil {
		// 找到指定节点 p，退出循环
		if current == p {
			break
		}
		previous = current     // 记录当前节点为前一个节点
		current = current.next // 移动到下一个节点
	}

	// 说明 p 为尾节点
	if current == nil {
		return false
	}

	// 删除节点 p，将前一个节点的 next 指针指向节点 p 的下一个节点
	previous.next = p.next

	list.length--
	return true
}

// 查找指定值的节点
func (list *LinkedList[T]) Find(v T) *Node[T] {
	current := list.head.next

	for current != nil {
		// 对比值是否匹配，如果匹配则返回该节点
		if current.GetValue().Compare(v) {
			return current
		}
		current = current.GetNext()
	}
	// 未找到匹配的节点，返回 nil
	return nil
}

func (list *LinkedList[T]) Update(oldValue, newValue T) bool {
	// 查找链表中的节点
	node := list.Find(oldValue)

	if node == nil {
		return false
	}
	// 更新节点的值
	node.data = newValue
	return true
}
