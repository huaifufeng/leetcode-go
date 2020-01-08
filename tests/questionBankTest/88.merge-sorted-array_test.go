package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestMerge(t *testing.T) {
	questionBank.Merge([]int{0}, 0, []int{1}, 1)
}