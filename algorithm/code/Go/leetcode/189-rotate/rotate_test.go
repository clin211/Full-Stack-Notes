package rotate

import (
	"reflect"
	"testing"
)

// 测试函数
func TestRotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2}, 1, []int{2, 1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},              // k 等于数组长度
		{[]int{}, 1, []int{}},                            // 空数组
		{[]int{1}, 1, []int{1}},                          // 单个元素
		{[]int{1, 2, 3, 4, 5}, 10, []int{1, 2, 3, 4, 5}}, // k 大于数组长度
	}

	for _, tt := range tests {
		// 深拷贝 nums，以防直接修改
		nums := make([]int, len(tt.nums))
		copy(nums, tt.nums)
		rotate(nums, tt.k)
		if !reflect.DeepEqual(nums, tt.want) {
			t.Errorf("rotate(%v, %d) = %v; want %v", tt.nums, tt.k, nums, tt.want)
		}
	}
}
