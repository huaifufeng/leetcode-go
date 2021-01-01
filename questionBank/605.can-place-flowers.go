/**
  题目地址：https://leetcode-cn.com/problems/can-place-flowers/

  假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
  给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数n。能否在不打破种植规则的情况下种入n朵花？能则返回True，不能则返回False。

  <pre>
     输入：flowerbed = [1,0,0,0,1], n = 1
     输出：True
  </pre>

  <pre>
     输入: flowerbed = [1,0,0,0,1], n = 2
     输出: False
  </pre>

  解题：
	1、依次判断每个为0值的位置，判断前后是否也为0，这样的位置可以填入，n减一，如果n小于等于0就返回true，否则返回false
    要判断边界值情况，只有一个值，并且为0；不填入值 这些情况

*/
package questionBank

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 && flowerbed[0] == 0 && n <= 1 {
		return true
	}

	length := len(flowerbed)
	for i := 0; i < length; i++ {
		if flowerbed[i] == 1 {
			continue
		}

		if i == 0 {
			if i+1 < length && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else if i == length-1 {
			if i-1 >= 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else {
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}

		if n <= 0 {
			return true
		}
	}

	if n <= 0 {
		return true
	}

	return false
}
