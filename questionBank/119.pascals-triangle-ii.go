/**
  题目地址：https://leetcode-cn.com/problems/pascals-triangle-ii/

  给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
  在杨辉三角中，每个数是它左上方和右上方的数的和。

  <pre>
    输入: 3
    输出: [1,3,3,1]
  </pre>

  解题：
	1、对三角的每一行进行处理，处理第一行时固定的1，其他的都是每行开始和结尾都是1，其他的是该元素左上和右上的元素之和 O(n^2)
*/
package questionBank

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		 return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	ret := make([]int, rowIndex + 1)
	lastRow := getRow(rowIndex - 1)
	for i := 0; i <= rowIndex; i++ {
		if i == 0 || i == rowIndex{
			ret[i] = 1
		} else {
			ret[i] = lastRow[i - 1] + lastRow[i]
		}
	}

	return ret
}

func GetRow(rowIndex int) []int {
	return getRow(rowIndex)
}