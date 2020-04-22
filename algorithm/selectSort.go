package algorithm

//selectSort 选择排序，每次从没有排序的部分中获取最小的元素，然后和当前选择的元素进行交换
func selectSort(nums []int) []int{
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}

	return nums
}
