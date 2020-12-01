package msort

var aux []int

// MergeSort
// 归并排序
// 这是一个分治算法, 将一个数组分成左右两部分, 分别对左右两部分排序, 再将其合并.
func MergeSort(nums []int) {
	aux = make([]int, len(nums)) // 一次性分配空间
	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	sort(nums, lo, mid)      // 左半边排序
	sort(nums, mid+1, hi)    // 右半边排序
	merge(nums, lo, mid, hi) // 归并结果
}

func merge(nums []int, lo, mid, hi int) {
	// 将nums[lo..mid] 和 nums[mid+1..hi]归并
	i, j := lo, mid+1

	// 将nums[lo..hi]复制到aux[lo..hi]
	for k := lo; k <= hi; k++ {
		aux[k] = nums[k]
	}

	// 归并回到nums[lo..hi]
	for k := lo; k <= hi; k++ {
		if i > mid {
			nums[k] = aux[j] // aux左半边用尽, 取右半边
			j++
		} else if j > hi {
			nums[k] = aux[i] // aux右半边用尽, 取左半边
			i++
		} else if aux[j] < aux[i] {
			nums[k] = aux[j] // aux右半边元素小于当前左半边元素, 取右半边
			j++
		} else {
			nums[k] = aux[i] // aux右半边元素大于当前左半边元素, 去左半边
			i++
		}
	}
}
