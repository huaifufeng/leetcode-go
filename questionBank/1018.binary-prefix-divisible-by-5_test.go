package questionBank

import "testing"

func TestPrefixesDivBy5(t *testing.T) {
	nums := []int{0, 1, 1}
	res := prefixesDivBy5(nums)
	if res[0] != true || res[1] != false || res[2] != false {
		t.Error("可被5整除的数计算错误1！")
	}

	nums = []int{1, 1, 1}
	res = prefixesDivBy5(nums)
	if res[0] != false || res[1] != false || res[2] != false {
		t.Error("可被5整除的数计算错误2！")
	}

	nums = []int{0, 1, 1, 1, 1, 1}
	res = prefixesDivBy5(nums)
	if res[0] != true || res[1] != false || res[2] != false || res[3] != false || res[4] != true || res[5] != false {
		t.Error("可被5整除的数计算错误3！")
	}

	nums = []int{1, 1, 1, 0, 1}
	res = prefixesDivBy5(nums)
	if res[0] != false || res[1] != false || res[2] != false || res[3] != false || res[4] != false {
		t.Error("可被5整除的数计算错误4！")
	}
}
