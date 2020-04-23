package algorithm

//mergeSortFunc 声明进行处理任务的递归方法
func MergeSort(nums []int, left, right int) {
	if left < right {
		mid := (left + right ) / 2
		MergeSort(nums, left, mid)
		MergeSort(nums, mid + 1, right)
		merge(nums, left, mid, right)
	}
}

//merge 对两个处理完毕的子任务进行合并操作的方法
func merge(nums []int, left, mid, right int){
	i := left
	j := mid + 1
	k := 0
	//需要声明临时数组来存放合并之后的数据，如果直接在原有数组上进行操作，数组的值就会被修改了，
	tempArr := make([]int, right - left + 1)
	for i <= mid && j <= right {
		if nums[i] > nums[j]  {
			tempArr[k] = nums[j]
			k++
			j++
		} else {
			tempArr[k] = nums[i]
			k++
			i++
		}
	}

	start := i
	end := mid
	//如果i>mid 说明，左侧的数据已经处理完毕了，这里需要处理的是右侧的数据
	if i > mid {
		start = j
		end = right
	}

	for ; start <= end; start++ {
		tempArr[k] = nums[start]
		k++
	}

	for i = 0; i < len(tempArr); i++ {
		nums[left] = tempArr[i]
		left++
	}
}

