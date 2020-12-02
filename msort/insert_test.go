package msort

import (
	"reflect"
	"testing"
)

func BenchmarkInsertSort(b *testing.B) {
	nums := []int{9, 7, 8, 2, 10}
	assert := []int{2, 7, 8, 9, 10}
	InsertSort(nums)
	if !reflect.DeepEqual(nums, assert) {
		b.Fatal("排序错误", nums, "\n")
	}
}
