/**
  题目地址：https://leetcode-cn.com/problems/two-sum/

  给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
  你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

  <pre>
     给定 nums = [2, 7, 11, 15], target = 9

     因为 nums[0] + nums[1] = 2 + 7 = 9
     所以返回 [0, 1]
  </pre>

  解题：
  1、最简单的方法就是循环整个数组，然后依次将数组中的值相加，然后得到想要的结果。但是这种在数量很大的时候，会非常耗费时间，时间复杂度O(n²)
  2、我们循环数组，将数组的值和索引以 值：索引的方式放入一个map中，这样在后面的循环中，只需要判断目标值-当期值=差值 这个差值在不在map中，
  只要在map中存在，就说明找到了目前的组合(因为这里确定有且只有一个组合存在，那么找到一个组合就完成任务了)，时间复杂度O(n)
*/
package questionBank

func twoSum(num []int, target int) []int {
	numLength := len(num)
	//声明值与索引对应关系的map，这里map最长为数组的长度
	numMap := make(map[int]int, numLength)
	retNum := make([]int, 2)

	for i := 0; i < numLength; i++ {
		lackValue := target - num[i]
		//如果差值存在，说明在之前的循环中已经处理了这个值
		if lackIndex, ok := numMap[lackValue]; ok {
			retNum[0] = lackIndex
			retNum[1] = i
			return retNum
		}

		numMap[num[i]] = i
	}

	return retNum
}

//测试用例使用的方法
func TwoSum(num []int, target int) []int {
	return twoSum(num, target)
}
