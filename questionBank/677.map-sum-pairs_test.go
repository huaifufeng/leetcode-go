package questionBank

import "testing"

func TestMapSumPairs(t *testing.T) {
	mapSum := ConstructorMapSum()
	mapSum.Insert("apple", 3)
	res := mapSum.Sum("ap")
	if res != 3 {
		t.Error("键值映射算法测试1错误")
	}
	mapSum.Insert("app", 2)
	res = mapSum.Sum("ap")
	if res != 5 {
		t.Error("键值映射算法测试2错误")
	}
}
