package questionBank

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	res := climbStairs(4)
	if res != 5 {
		t.Error("爬楼算法错误")
	}
}
