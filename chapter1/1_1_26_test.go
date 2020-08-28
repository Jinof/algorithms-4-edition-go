package chapter1

import (
	"testing"
)

func TestSortABC(t *testing.T) {
	var (
		a, b, c                 int
		sorted, measure, source []int
	)

	a, b, c = 1, 10, 100
	measure = []int{1, 10, 100}

	// 测试已排序a, b, c
	source = []int{a, b, c}

	sorted = SortABC(source)
	if !equalABC(sorted, measure) {
		t.Fatalf("err: %s \n source: %v \n sorted: %v \n measure: %v ", "error in SortABC", source, sorted, measure)
	}

	// 测试反向数据
	a, b, c = 100, 10, 1
	source = []int{a, b, c}

	sorted = SortABC(source)
	if !equalABC(sorted, measure) {
		t.Fatalf("err: %s \n source: %v \n sorted: %v \n measure: %v ", "error in SortABC", source, sorted, measure)
	}
	t.Logf("%s", "原函数可以将abc按照升序排列")
}

func equalABC(sorted, measure []int) bool {
	for i, v := range sorted {
		if v != measure[i] {
			return false
		}
	}

	return true
}
