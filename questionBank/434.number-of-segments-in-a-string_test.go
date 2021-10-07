package questionBank

import "testing"

func TestCountSegments(t *testing.T) {
	str := "Hello, my name is John"
	res := countSegments(str)
	if res != 5 {
		t.Error("字符串中的单词数算法测试1错误", res)
	}

	str = " Hello, my  name is John "
	res = countSegments(str)
	if res != 5 {
		t.Error("字符串中的单词数算法测试2错误", res)
	}
}
