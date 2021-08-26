package questionBank

import "testing"

func TestNumRescueBoats(t *testing.T) {
	people := []int{1, 2}
	res := numRescueBoats(people, 3)
	if res != 1 {
		t.Error("救生艇算法测试1错误！")
	}

	people = []int{3, 2, 2, 1}
	res = numRescueBoats(people, 3)
	if res != 3 {
		t.Error("救生艇算法测试2错误！")
	}

	people = []int{3, 5, 3, 4}
	res = numRescueBoats(people, 5)
	if res != 4 {
		t.Error("救生艇算法测试2错误！")
	}
}
