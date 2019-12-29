package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	res := questionBank.RemoveDuplicates(nums)
	if res != 5 {
		t.Log(res)
		t.Error("方法验证失败")
	}

	//因为传入的是切片，所以在方法里面做的修改，都可以在变量上面展示
	//下面的方法可以展示修改之后的不重复元素值
	t.Error(nums[:res])
}