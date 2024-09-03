package productexceptself

import (
	"reflect"
	"testing"
)

// 测试函数
func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{0, 0}, []int{0, 0}},          // 两个零的情况
		{[]int{1, 0, 3}, []int{0, 3, 0}},    // 一个零的情况
		{[]int{5, 6, 2}, []int{12, 10, 30}}, // 正数情况
		{[]int{-1, -2, -3}, []int{6, 3, 2}}, // 负数情况
	}

	for _, tt := range tests {
		got := productExceptSelf(tt.nums)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("productExceptSelf(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}
