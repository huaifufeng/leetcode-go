package algorithm

//插入排序
//稳定排序，时间复杂度O(n^2) 空间复杂度O(1)  原地排序

//InsertSort 插入排序，采用从后到前的顺序，依次找到i位置的元素在前面有序列表中的合适位置，放入进去
func InsertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				break
			}
		}
	}
}

//InsertSort2 相比于上面的方法，这个方法保存了当前元素的值，没有只需要吧前面的值往后面移动就可以，减少了交换的次数
func InsertSort2(nums []int) {
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > curr {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}

		nums[j+1] = curr
	}
}
