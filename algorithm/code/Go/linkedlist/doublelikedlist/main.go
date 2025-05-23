package main

import "fmt"

// 定义双向链表的节点结构体
type Node struct {
	Data int
	prev *Node // 指向前一个节点
	next *Node // 指向后一个节点
}

// 定义双向链表结构体
type DoublyLinkedList struct {
	head *Node // 指向链表头部
	tail *Node // 指向链表尾部
}

// 创建一个新的节点
func newNode(data int) *Node {
	return &Node{Data: data}
}

// 在双向链表的头部插入一个新节点
func (dll *DoublyLinkedList) InsertAtHead(data int) {
	node := newNode(data)
	if dll.head == nil {
		dll.head = node
		dll.tail = node
		return
	}
	dll.head.next = node
	node.prev = dll.tail
	dll.tail = node
}

// 在双向链表的指定位置插入一个新节点
func (dll *DoublyLinkedList) insertAt(data int, position int) {
	if position < 0 {
		// 负数位置无效
		return
	}

	newNode := newNode(data) // 创建新节点
	if dll.head == nil {
		// 如果链表为空，新节点成为唯一的节点
		dll.head = newNode
		dll.tail = newNode
		return
	}

	current := dll.head
	for current != nil && position > 0 {
		// 移动到指定位置
		current = current.next
		position--
	}

	if current == nil {
		// 插入位置超出链表长度，将新节点插入到末尾
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	} else {
		// 插入新节点到链表中
		newNode.next = current
		newNode.prev = current.prev
		if current.prev != nil {
			current.prev.next = newNode
		}
		current.prev = newNode
	}
}

// 在双向链表的末尾插入一个新节点
func (dll *DoublyLinkedList) append(data int) {
	newNode := newNode(data)
	if dll.tail == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}
	dll.tail.next = newNode
	newNode.prev = dll.tail
	dll.tail = newNode
}

// 删除双向链表中的一个节点
func (dll *DoublyLinkedList) remove(node *Node) {
	if node == nil {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		dll.head = node.next // 删除头节点
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		dll.tail = node.prev // 删除尾节点
	}
}

// 查找双向链表中的一个节点
func (dll *DoublyLinkedList) find(data int) *Node {
	current := dll.head
	for current != nil {
		if current.Data == data {
			return current
		}
		current = current.next
	}
	return nil
}

// 修改节点的数据
func (dll *DoublyLinkedList) update(oldData, newData int) bool {
	node := dll.find(oldData)
	if node != nil {
		node.Data = newData
		return true
	}
	return false
}

// 遍历双向链表
func (dll *DoublyLinkedList) traverse() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	dll := DoublyLinkedList{}

	// 插入元素
	dll.append(1)
	dll.append(2)
	dll.append(3)

	// 遍历链表
	fmt.Println("Traverse forward:")
	dll.traverse()

	// 查找元素
	node := dll.find(2)
	if node != nil {
		fmt.Printf("Found node with data: %d\n", node.Data)
	}

	// 删除元素
	dll.remove(dll.find(2))

	// 再次遍历链表
	fmt.Println("Traverse forward after removal:")
	dll.traverse()

	// 反向遍历链表
	fmt.Println("Traverse backward:")
	current := dll.tail
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.prev
	}
	fmt.Println()
}
