package algorithm

import (
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArray(5)
	res := arr.insert(0, 1)
	if res != true || arr.count != 1 {
		t.Error("数组的插入方法错误")
	}
	res = arr.insert(1, 2)
	if res != true || arr.count != 2 {
		t.Error("数组的插入方法错误")
	}
	res = arr.insert(3, 4)
	if res != false || arr.count != 2 {
		t.Error("数组的插入方法错误")
	}
	res = arr.insert(2, 3)
	if res != true || arr.count != 3 {
		t.Error("数组的插入方法错误")
	}

	res = arr.insert(1, 5)
	if res != true || arr.count != 4 || arr.data[1] != 5 {
		t.Error("数组的插入方法错误")
	}

	res = arr.delete(1)
	if res != true || arr.count != 3 || arr.data[1] != 2 {
		t.Error("数组的插入方法错误")
	}
}
