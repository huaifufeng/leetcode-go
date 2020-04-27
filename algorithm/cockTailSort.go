package algorithm

//鸡尾酒排序
//解决冒泡排序的全逆序的问题 先右序排序 在左序排序
func CockTailSort(nums []int) {
	//因为每次排序里面有两次排序，外部只需要进行一半的排序处理就可以了
	for i := 0; i < len(nums)/2; i++ {
		isSorted := false
		//第一遍从左往右处理 这里j=i 是因为每次循环处理之后，前后各有一个已经处理完毕
		for j := i; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSorted = true
			}
		}

		if !isSorted {
			break
		}

		isSorted = false
		//从右向左排序 取值范围是每次处理之后 都会前后各有一个处理完毕
		for j := len(nums) - i - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				isSorted = true
			}
		}

		if !isSorted {
			break
		}
	}

}

//增加排序区间的处理
func CockTailSort2(nums []int) {
	leftSortIndex := 0
	rightSortIndex := len(nums) - 1
	lastIndex := 0
	for i := 0; i < len(nums)/2; i++ {
		isSorted := true
		//第一次交换的最后交换位置 确认了右交换位的索引
		for j := leftSortIndex; j < rightSortIndex; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSorted = false
				lastIndex = j
			}
		}

		rightSortIndex = lastIndex
		if isSorted {
			break
		}

		isSorted = true
		//第二次交换的最后交换位置，确认了左交换位的索引
		for j := rightSortIndex; j > leftSortIndex; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				isSorted = false
				lastIndex = j
			}
		}

		leftSortIndex = lastIndex
		if isSorted {
			break
		}
	}

}
