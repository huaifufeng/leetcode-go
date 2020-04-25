package algorithm

//冒泡排序
//稳定排序，时间复杂度O(n^2) 空间复杂度O(1)  原地排序

//BubbleSort 冒泡排序算法，每个位置和临位进行比较。增加一个标识为，如果没有进行交换处理，说明这些元素不用处理直接结束
func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		//标识，判断是否有没有交换了，这时就表示都已经排序好了
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				//有交换行为，说明还存在没有排好序的数据
				flag = true
			}
		}

		if !flag {
			break
		}
	}
}

//BubbleSort2 相比于上面的方法，增加交换位标识，交换位之后的位置，没有交换处理，说明都是已经排好序的元素
func BubbleSort2(nums []int) {
	sortPos := len(nums) - 1
	lastIndex := 0
	for i := 0; i < len(nums); i++ {
		sorted := false
		for j := 0; j < sortPos; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				sorted = true
				lastIndex = j
			}
		}

		sortPos = lastIndex
		if !sorted {
			break
		}
	}
}
