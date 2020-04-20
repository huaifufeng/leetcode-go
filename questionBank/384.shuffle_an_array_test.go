package questionBank

import (
	"testing"
)

func TestShuffleAnArray(t *testing.T) {
	nums := []int{1,2,3}
	_, p2 := ShuffleAnArray(nums)

	for key, value := range p2 {
		if nums[key] != value {
			t.Errorf("重置方法出现错误")
		}
	}
}
