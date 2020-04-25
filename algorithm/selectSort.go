package algorithm

//选择排序
//不稳定排序，时间复杂度O(n^2) 空间复杂度O(1) 原地排序

//selectSort 选择排序，每次从没有排序的部分中获取最小的元素，然后和当前选择的元素进行交换
func SelectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		if i != minIndex {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
}
