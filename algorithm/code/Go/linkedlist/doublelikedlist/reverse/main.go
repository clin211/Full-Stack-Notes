package main

// 定义 Node
type Node struct {
	data int
	prev *Node
	next *Node
}

type List struct {
	head   *Node
	tail   *Node
	length uint
}

func (list *List) reverse() {
	if list.length <= 0 {
		return
	}

	var prev *Node
	current := list.head
	list.tail = list.head // 反转后，原头节点变成尾节点
	for current != nil {
		// 先保存当前节点的下一个节点
		next := current.next
		// 反转当前节点的 next 节点
		current.next = prev
		// 反正当前节点的 prev 节点
		current.prev = next
		// 移动 prev 指针
		prev = current
		// 移动下一个节点
		current = next
	}
	list.head = prev
}
