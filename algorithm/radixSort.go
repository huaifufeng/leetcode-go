package algorithm

//基数排序
//稳定排序 时间复杂度O(n) 空间复杂度O(n)

func RadixSort(nums []string) {
	//获取数组中最长的字符串的长度
	strLength := 0
	for i := 0; i < len(nums); i++ {
		if strLength < len(nums[i]) {
			strLength = len(nums[i])
		}
	}

	for i := strLength - 1; i >= 0; i-- {
		//首先获取进行计数排序的数组
		sortArray := make([]byte, len(nums))
		max := nums[0][i]
		min := nums[0][i]
		for j := 0; j < len(nums); j++ {
			sortArray[j] = getChar(nums[j], i)
			if sortArray[j] > max {
				max = sortArray[j]
			} else if sortArray[j] < min {
				min = sortArray[j]
			}
		}

		//获取每个元素的数量
		countArray := make([]byte, max-min+1)
		for j := 0; j < len(nums); j++ {
			index := getChar(nums[j], i) - min
			countArray[index]++
		}
		//获取每个元素的位置
		for j := 1; j < len(countArray); j++ {
			countArray[j] += countArray[j-1]
		}

		sortedArray := make([]string, len(nums))
		for j := len(nums) - 1; j >= 0; j-- {
			countIndex := getChar(nums[j], i) - min
			index := countArray[countIndex] - 1
			sortedArray[index] = nums[j]
			countArray[countIndex]--
		}

		copy(nums, sortedArray)
	}
}

func getChar(str string, k int) byte {
	if len(str) < k+1 {
		return '0'
	}

	return str[k]
}
