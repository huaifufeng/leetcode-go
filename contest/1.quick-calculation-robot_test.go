package contest

import "testing"

func TestCalculate(t *testing.T) {
	str := "AB"
	res := calculate(str)
	if res != 4 {
		t.Error("速算机器人算法错误！")
	}
}
