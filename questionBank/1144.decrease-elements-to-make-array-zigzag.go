/**
  题目地址：https://leetcode-cn.com/problems/decrease-elements-to-make-array-zigzag/

  给你一个整数数组nums，每次 操作会从中选择一个元素并 将该元素的值减少1。
  如果符合下列情况之一，则数组A就是 锯齿数组：
    1、每个偶数索引对应的元素都大于相邻的元素，即A[0] > A[1] < A[2] > A[3] < A[4] > ...
    2、或者，每个奇数索引对应的元素都大于相邻的元素，即A[0] < A[1] > A[2] < A[3] > A[4] < ...
    返回将数组nums转换为锯齿数组所需的最小操作次数。

  <pre>
    输入：nums = [1,2,3]
    输出：2
    解释：我们可以把 2 递减到 0，或把 3 递减到 1。
  </pre>

  解题：
     1、有两种生成方法，分别计算两种生成方式，其中最小的值就是我们需要的结果

  PS:
    1、需要注意处理的元素结果需要放入数组，这样避免了重复计算
    2、对切片的操作是会记录的，这里需要复制一份切片数据出来
*/
package questionBank

import "fmt"

func movesToMakeZigzag(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	//初始化两种方式的执行次数
	ordNum := 0
	evenNum := 0

	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	//根据元素的位置分为两种情况进行处理：
	//1、对第二种情况进行处理，只处理奇数的值 增加evenNum的值
	for index := range copyNums {
		if index%2 == 1 {
			leftSub := 0
			rightSub := 0
			//不处理小于0的位置的元素
			leftSub = copyNums[index-1] - copyNums[index] + 1
			if leftSub > 0 {
				copyNums[index-1] -= leftSub
				evenNum += leftSub
			}

			//不处理超出范围的元素
			if index+1 < len(copyNums) {
				rightSub = copyNums[index+1] - copyNums[index] + 1
				if rightSub > 0 {
					copyNums[index+1] -= rightSub
					evenNum += rightSub
				}
			}
		}
	}

	//2、对第一种方式进行处理
	for index := range nums {
		if index%2 == 0 {
			leftSub := 0
			rightSub := 0
			//不处理小于0的位置的元素
			if index-1 >= 0 {
				leftSub = nums[index-1] - nums[index] + 1
				if leftSub > 0 {
					nums[index-1] -= leftSub
					ordNum += leftSub
				}
			}

			//不处理超出范围的元素
			if index+1 < len(nums) {
				rightSub = nums[index+1] - nums[index] + 1
				fmt.Println(index, nums)
				if rightSub > 0 {
					nums[index+1] -= rightSub
					ordNum += rightSub
				}
			}
		}
	}

	fmt.Println(ordNum, evenNum)
	minNum := ordNum
	if minNum > evenNum {
		minNum = evenNum
	}

	return minNum
}
