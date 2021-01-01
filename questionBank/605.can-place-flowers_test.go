package questionBank

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	flowerbed := []int{0}

	res := canPlaceFlowers(flowerbed, 1)
	if res != true {
		t.Error("是否可以放置花瓶出现错误")
	}

	res = canPlaceFlowers(flowerbed, 2)
	if res != false {
		t.Error("是否可以放置花瓶出现错误")
	}
}
