package algorithm

//二分查找
//时间复杂度O(logn)  空间复杂度O(1)
func BinarySearch(nums []int, needle int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		//避免直接相加超出数据类型范围，这里使用加上插值的方法
		//可以使用位移，性能更好 left + ((right-left)>>1)
		mid := left + (right-left)/2
		if nums[mid] < needle {
			left = mid + 1
		} else if nums[mid] > needle {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

//以下为数组有序，其中有重复元素存在

//获取元素在数组中第一次出现的位置
func BinarySearchFirst(nums []int, needle int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > needle {
			right = mid - 1
		} else if nums[mid] < needle {
			left = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != needle {
				return mid
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

//获取元素在数组中最后出现的位置
func BinarySearchLast(nums []int, needle int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > needle {
			right = mid - 1
		} else if nums[mid] < needle {
			left = mid + 1
		} else {
			if mid == right || nums[mid+1] != needle {
				return mid
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

//查找第一个大于等于元素的位置
func BinarySearchFirstLarge(nums []int, needle int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < needle {
			left = mid + 1
		} else {
			if mid == 0 || nums[mid-1] < needle {
				return mid
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

//查找最后一个小于等于元素的值
func BinarySearchLastMin(nums []int, needle int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > needle {
			right = mid - 1
		} else {
			if mid == right || nums[mid+1] > needle {
				return mid
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
