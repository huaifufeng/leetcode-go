/**
  题目地址：https://leetcode-cn.com/problems/sliding-window-maximum/

  给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
  返回滑动窗口中的最大值。

  <pre>
    输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
	输出：[3,3,5,5,6,7]
	解释：
	滑动窗口的位置                最大值
	---------------               -----
	[1  3  -1] -3  5  3  6  7       3
	 1 [3  -1  -3] 5  3  6  7       3
	 1  3 [-1  -3  5] 3  6  7       5
 	 1  3  -1 [-3  5  3] 6  7       5
 	 1  3  -1  -3 [5  3  6] 7       6
 	 1  3  -1  -3  5 [3  6  7]      7
  </pre>

  <pre>
    输入：nums = [1,-1], k = 1
	输出：[1,-1]
  </pre>

  解题：
    1、这里是从左往右依次处理每个分组块，但是为了避免超时的情况，对于新加入的值和之前的大值进行比较，如果小于之前的大值，并且之前的大值不是移除的，就直接放入之前的最大值，避免一次分组比较
*/
package questionBank

func maxSlidingWindow(nums []int, k int) []int {
	//当k=1时，表示元素每个值都是最大的
	if k == 1 {
		return nums
	}

	var maxNums []int
	lastIndex := 0
	maxNum := -10000
	endIndex := len(nums) - k

	for lastIndex <= endIndex {
		if lastIndex != 0 {
			newNum := nums[lastIndex+k-1]
			if newNum < maxNum && (nums[lastIndex-1] != maxNum || nums[lastIndex] == maxNum) {
				maxNums = append(maxNums, maxNum)
			} else {
				subNums := nums[lastIndex : lastIndex+k]
				maxNum = getMaxNum(subNums)
				maxNums = append(maxNums, maxNum)
			}
		} else {
			subNums := nums[lastIndex : lastIndex+k]
			maxNum = getMaxNum(subNums)
			maxNums = append(maxNums, maxNum)
		}

		lastIndex++
	}

	return maxNums
}

func getMaxNum(nums []int) int {
	maxNum := -10000
	for _, val := range nums {
		if val > maxNum {
			maxNum = val
		}
	}

	return maxNum
}
