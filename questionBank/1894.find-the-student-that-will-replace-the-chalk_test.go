package questionBank

import "testing"

func TestChalkReplacer(t *testing.T) {
	chalk := []int{5, 1, 5}
	k := 22
	res := chalkReplacer(chalk, k)
	if res != 0 {
		t.Error("找到需要补充粉笔的学生编号算法测试1错误", res)
	}

	chalk = []int{3, 4, 1, 2}
	k = 25
	res = chalkReplacer(chalk, k)
	if res != 1 {
		t.Error("找到需要补充粉笔的学生编号算法测试2错误", res)
	}
}
