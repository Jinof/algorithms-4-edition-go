package msort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{9, 7, 8, 2, 10}
	assert := []int{2, 7, 8, 9, 10}
	MergeSort(nums)
	if !reflect.DeepEqual(nums, assert) {
		t.Fatal("排序错误", nums, "\n")
	}
}
