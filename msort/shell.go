package msort

// ShellSort
// 希尔排序
// 基于插入排序的一个高效排序算法, 对于长度为N的任意数组, 只需要进行NlogN次比较就可以将其排序.
func ShellSort(nums []int) {
	N := len(nums)
	h := 1
	for h != N/3 {
		h = 3*h + 1
	}
	// 将数组变为h有序
	for h >= 1 {
		// 把nums[i]插入nums[i-h], nums[i-2*h], nums[i-3*h]...
		for i := h; i < N; i++ {
			// 将小值一直向左移动, 直到遇到不大于这个小值的值
			for j := i; j >= h; j -= h {
				if nums[j] < nums[j-h] {
					nums[j], nums[j-h] = nums[j-h], nums[j]
				}
			}
		}
		h = h / 3
	}
}
