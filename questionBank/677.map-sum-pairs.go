/**
  题目地址：https://leetcode-cn.com/problems/map-sum-pairs/

  解题：
    1、循环1~n，判断每个数字和3，5的关系，防入数组 O(n)
*/
package questionBank

import "strings"

type MapSum struct {
	Dict map[string]int
}

func ConstructorMapSum() MapSum {
	var mapSum = MapSum{
		Dict: make(map[string]int),
	}

	return mapSum
}

func (this *MapSum) Insert(key string, val int) {
	this.Dict[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	total := 0
	for k, v := range this.Dict {
		if strings.HasPrefix(k, prefix) {
			total += v
		}
	}

	return total
}
