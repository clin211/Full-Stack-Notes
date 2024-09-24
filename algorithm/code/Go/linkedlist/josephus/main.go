package main

// 约瑟夫环问题
// 约瑟夫环（约瑟夫斯问题）是一个数学应用问题：
// 已知 n 个人（以编号 1，2，3，...，n 分别表示）围坐在一张圆桌周围，
// 从编号为 k 的人开始报数，数到 m 的那个人出列；
// 他的下一个人又从 1 开始报数，数到 m 的那个人又出列；
// 依此规律重复下去，直到圆桌周围只剩一个人。
// 通常解决这个问题的步骤如下：
// 1. 创建一个包含 n 个节点的循环链表。
// 2. 从第 k 个节点开始，每次数到 m 的那个节点出列。
// 3. 重复步骤 2，直到链表中只剩下一个节点。
// 4. 返回这个节点的编号。
type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
}

func (cll *CircularLinkedList) Insert(data int) {
	newNode := &Node{data: data}
	if cll.head == nil {
		cll.head = newNode
		newNode.next = cll.head
	} else {
		current := cll.head
		for current.next != cll.head {
			current = current.next
		}
		current.next = newNode
		newNode.next = cll.head
	}
}

func (cll *CircularLinkedList) Delete(node *Node) {
	if cll.head == nil {
		return
	}
	if cll.head == node && cll.head.next == cll.head {
		cll.head = nil
		return
	}
	current := cll.head
	for current.next != cll.head {
		if current.next == node {
			current.next = current.next.next
			if node == cll.head {
				cll.head = current.next
			}
			return
		}
		current = current.next
	}
}

// Josephus 函数用于解决约瑟夫环问题
// n 是总人数，k 是开始报数的人的编号，m 是每次报数的间隔
func Josephus(n, k, m int) int {
	cll := &CircularLinkedList{}
	for i := 1; i <= n; i++ {
		cll.Insert(i)
	}

	current := cll.head
	for i := 1; i < k; i++ {
		current = current.next
	}

	for cll.head.next != cll.head {
		for i := 1; i < m; i++ {
			current = current.next
		}
		next := current.next
		cll.Delete(current)
		current = next
	}

	return cll.head.data
}

func main() {

}
