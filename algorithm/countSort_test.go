package algorithm

import "testing"

func TestCountSort(t *testing.T) {
	nums := []int{1, 3, 2, 7, 3, 5, 1, 4}
	sortedNums := CountSort(nums)

	for i := 0; i < len(sortedNums)-1; i++ {
		if sortedNums[i] > sortedNums[i+1] {
			t.Error("计数排序结果有问题")
		}
	}
}
