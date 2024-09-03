package rotate

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	// 处理 k 大于数组长度的情况
	k = k % n
	if k == 0 {
		return
	}

	// 反转整个数组
	reverse(nums, 0, n-1)
	// 反转前 k 个元素
	reverse(nums, 0, k-1)
	// 反转后 n-k 个元素
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		// 交换元素
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
