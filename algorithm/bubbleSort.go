package algorithm

func BubbleSort(nums []int) []int{
	if len(nums) <= 1 {
		return nums
	}

	for i := 0 ; i < len(nums); i++ {
		//标识，判断是否有没有交换了，这时就表示都已经排序好了
		flag := false
		for j := 0; j < len(nums) - i - 1; j++ {
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

	return nums
}

