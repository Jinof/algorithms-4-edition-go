package chapter1

//将三个数字排序。假设 a、b、c 和 t 都是同一种原始数字类型的变量。证明以下代码能够将 a、
//b、c 按照升序排列：
//if (a > b) { t = a; a = b; b = t; }
//if (a > c) { t = a; a = c; c = t; }
//if (b > c) { t = b; b = c; c = t; }

func SortABC(nums []int) []int {
	var (
		t       int
		a, b, c = nums[0], nums[1], nums[2]
	)

	if a > b {
		t = a
		a = b
		b = t
	}
	if a > c {
		t = a
		a = c
		c = t
	}
	if b > c {
		t = b
		b = c
		c = t
	}

	nums[0], nums[1], nums[2] = a, b, c
	return nums
}
