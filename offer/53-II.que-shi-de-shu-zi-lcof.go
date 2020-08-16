/**
  题目地址：https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

  一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
  在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

  <pre>
     输入: [0,1,3]
     输出: 2
  </pre>

  解题：
    1、暴力破解法。循环数组，比较数组的索引和元素值，不想等，说明就是当前索引缺失，否则最后一个元素缺失 O(n)
    2、二分查找。生成一个0~n-2的n-1位数组，每次比较中间索引位置，根据值的比较进行下一步的判断,不想等，说明就是当前索引缺失，否则最后一个元素缺失
*/
package offer

func missingNumber(nums []int) int {
	for index, value := range nums {
		if index != value {
			return index
		}
	}

	return len(nums)
}
