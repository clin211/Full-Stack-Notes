package main

// 递归解法
func josephusRecursive(n, m int) int {
	if n == 1 {
		return 0 // 只有一个人时，最后的幸存者编号为 0
	}
	return (josephusRecursive(n-1, m) + m) % n
}
