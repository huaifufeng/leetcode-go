package algorithm

//桶排序
//稳定排序 时间复杂度 O(n)  空间复杂度O(n)

func BucketSort(nums []int, bucketSize int) {
	max := nums[0]
	min := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
	}

	bucketNum := (max-min)/bucketSize + 1
	bucketArray := make([][]int, bucketNum)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < bucketNum; j++ {
			left := min + j*bucketSize
			right := min + (j+1)*bucketSize
			if nums[i] >= left && nums[i] < right {
				bucketArray[j] = append(bucketArray[j], nums[i])
				break
			}
		}
	}

	for i := 0; i < len(bucketArray); i++ {
		QuickSort(bucketArray[i], 0, len(bucketArray[i])-1)
	}

	index := 0
	for i := 0; i < len(bucketArray); i++ {
		for j := 0; j < len(bucketArray[i]); j++ {
			nums[index] = bucketArray[i][j]
			index++
		}
	}
}
