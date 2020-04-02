package questionBank

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	res := ClimbStairs(4)
	if res != 5 {
		t.Log(res)
		t.Error("爬楼算法错误")
	}
}