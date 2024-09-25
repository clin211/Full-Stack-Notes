package main

import (
	"fmt"
)

// n 参与人数，m 报数到 m 的人出局
func Josephus(n, m int) int {
	// 创建一个切片，用于存储参与者的编号（从 0 到 n-1）
	participants := make([]int, 0, n) // 预先分配 n 的容量，但初始长度为 0
	for i := 0; i < n; i++ {
		participants = append(participants, i) // 将每个参与者的编号添加到切片中
	}

	current := 0 // 初始化当前索引，从第一个参与者开始

	// 循环直到只剩下一个参与者
	for len(participants) > 1 {
		// 计算出局者的索引
		current = (current + m - 1) % len(participants)

		// 移除出局者
		participants = append(participants[:current], participants[current+1:]...)
	}

	return participants[0] // 返回最后剩下的参与者的编号
}

func main() {
	n := 7 // 参与者人数
	m := 3 // 报数到 m 的人出局

	lastPerson := Josephus(n, m)               // 调用函数计算最后剩下的人
	fmt.Printf("最后剩下的人的下标是: %d\n", lastPerson) // 输出结果
}
