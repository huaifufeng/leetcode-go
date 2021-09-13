/**
  题目地址：https://leetcode-cn.com/problems/number-of-boomerangs/

  解题：
  1、三层循环这个数组，依次计算第一层点到另外两个点的距离，相等的话加2，因为可以调整顺序，O(n^3)
  2、
*/
package questionBank

func numberOfBoomerangs(points [][]int) int {
	if len(points) < 3 {
		return 0
	}
	res := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}

			left := (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			for k := j + 1; k < len(points); k++ {
				right := (points[i][0]-points[k][0])*(points[i][0]-points[k][0]) + (points[i][1]-points[k][1])*(points[i][1]-points[k][1])
				if left == right {
					res += 2
				}
			}
		}
	}

	return res
}

func numberOfBoomerangs1(points [][]int) int {
	if len(points) < 3 {
		return 0
	}
	total := 0
	for _, point := range points {
		lengthMap := make(map[int]int, len(points))
		for _, ints := range points {
			length := (point[0]-ints[0])*(point[0]-ints[0]) + (point[1]-ints[1])*(point[1]-ints[1])
			lengthMap[length]++
		}

		for _, i2 := range lengthMap {
			total += i2 * (i2 - 1)
		}
	}

	return total
}
