/**
  题目地址：https://leetcode-cn.com/problems/pascals-triangle/

  给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
  在杨辉三角中，每个数是它左上方和右上方的数的和。
  说明：
    初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
    你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

  <pre>
    输入: 5
    输出:
    [
         [1],
        [1,1],
       [1,2,1],
      [1,3,3,1],
     [1,4,6,4,1]
    ]
  </pre>

  解题：
	1、对三角的每一行进行处理，处理第一行时固定的1，其他的都是每行开始和结尾都是1，其他的是该元素左上和右上的元素之和 O(n^2)
*/
package questionBank

func generate(numRows int) [][]int {
	//初始化要返回的切片的列表
	ret := make([][]int, numRows)

	//依次为每一行添加一个整数切片，当第一行时就是i=0时，只有1 一个值
	for i := 0; i < numRows; i++ {
		if i == 0 {
			ret[i] = []int{1}
			continue
		}

		ret[i] = make([]int, i + 1)
		//当第一行之外的行数时，循环添加i+1个元素，其中0和i位的位置为1，其他位置为上一行中左上元素和右上元素的和
		for j := 0; j <= i; j++ {
			lastLine := i - 1
			//一行还是和结尾的元素为1
			if j == 0 || j == i {
				ret[i][j] = 1
				continue
			} else {
				ret[i][j] = ret[lastLine][j-1] + ret[lastLine][j]
			}
		}
	}

	return ret
}

func Generate(numRows int) [][]int {
	return generate(numRows)
}