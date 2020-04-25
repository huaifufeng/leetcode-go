package algorithm

//计数排序
//稳定排序，时间复杂度O(n+k) 空间复杂度O(n)
//条件：整数，小数量，间隔不大

func CountSort(nums []int) []int {
	min := nums[0]
	max := nums[0]

	//获取数组的最大和最小值， 进行计数数组的长度计算
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
	}

	length := max - min + 1
	countArray := make([]int, length)
	//汇总出每个位置元素的数量
	for i := 0; i < len(nums); i++ {
		index := nums[i] - min
		countArray[index]++
	}

	//将数组的元素从左向右加起来，获取计数数组索引值的位置
	for i := 1; i < len(countArray); i++ {
		countArray[i] += countArray[i-1]
	}

	sortedArray := make([]int, len(nums))

	//获取每个元素的合适位置
	for i := len(nums) - 1; i >= 0; i-- {
		//元素位置的计算
		countIndex := nums[i] - min
		index := countArray[countIndex] - 1
		//这里需要减1 因为索引从0开始
		sortedArray[index] = nums[i]
		countArray[countIndex]--
	}

	//这里不是在原数组进行操作，所以需要返回这个数组
	return sortedArray
}
