package msort

// selectSort
// 选择排序
// 经典排序算法. 遍历数组, 选择最小值j, 交换j与i(i为当前循环所指位置)
// 遍历一次后即可将
func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if min > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
