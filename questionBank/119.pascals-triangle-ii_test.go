package questionBank

import (
	"testing"
)

func TestGetRow(t *testing.T) {
	res := getRow(3)
	if res[2] != 3 {
		t.Error("获取杨辉三角一行的方法错误")
	}
}
