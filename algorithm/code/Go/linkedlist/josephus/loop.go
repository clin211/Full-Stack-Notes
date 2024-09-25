package main

func josephusLoop(n, m int) int {
	position := 0              // 初始化位置为 0，表示只有一个人时的幸存者位置
	for i := 2; i < n+1; i++ { // 从 2 个人循环到 n 个人
		position = (position + m) % i // 计算当前 i 个人的幸存者位置
	}
	return position // 返回最后的幸存者位置
}
