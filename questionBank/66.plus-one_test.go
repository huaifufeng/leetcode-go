package questionBank

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	res := plusOne([]int{1, 2, 3})
	if res[2] != 4 {
		t.Error("数组整数加1出错")
	}
}
