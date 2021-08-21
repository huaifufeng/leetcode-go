package questionBank

import "testing"

func TestCompress(t *testing.T) {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	length := compress(chars)
	if length != 6 {
		t.Error("字符串压缩方法测试1失败", length)
	}

	chars = []byte{'a'}
	length = compress(chars)
	if length != 1 {
		t.Error("字符串压缩方法测试2失败", length)
	}

	chars = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	length = compress(chars)
	if length != 4 {
		t.Error("字符串压缩方法测试3失败", length)
	}
}
