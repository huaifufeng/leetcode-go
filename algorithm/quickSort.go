package algorithm

//快速排序
//不稳定排序，时间复杂度O(nlogn) 空间复杂度O(logn) 原地排序

func QuickSort(nums []int, start, end int) {
	if start <= end {
		pivot := partition2(nums, start, end)

		QuickSort(nums, start, pivot-1)
		QuickSort(nums, pivot+1, end)
	}
}

//partition 选择数组尾部节点作为基点
func partition(nums []int, start, end int) int {
	i := start
	j := start
	for ; j < end; j++ {
		if nums[j] <= nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[end] = nums[end], nums[i]

	return i
}

//partition2 选择排头作为基点
func partition2(nums []int, start, end int) int {
	i := end
	j := end
	for ; j > start; j-- {
		if nums[j] > nums[start] {
			nums[i], nums[j] = nums[j], nums[i]
			i--
		}
	}

	nums[i], nums[start] = nums[start], nums[i]

	return i
}

//快排优化
//1、 首中尾 各取一个，取中值，  大数量级可以 五取一 十取一
//2、 随机取值

//栈溢出
//1、限制栈的层数   2、自己模拟栈
