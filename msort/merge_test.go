package msort

import (
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	var nums []int
	copy(nums, NUMS)
	InsertSort(nums)
	if !sorted(nums) {
		b.Fatal("排序错误", nums, "\n")
	}
}
