/**
  题目地址：https://leetcode-cn.com/problems/shuffle-an-array/

  打乱一个没有重复元素的数组。

  <pre>
    // 以数字集合 1, 2 和 3 初始化数组。
	int[] nums = {1,2,3};
	Solution solution = new Solution(nums);

	// 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
	solution.shuffle();

	// 重设数组到它的初始状态[1,2,3]。
	solution.reset();

	// 随机返回数组[1,2,3]打乱后的结果。
	solution.shuffle();
  </pre>

  解题： 时间复杂度O(n) n为数组长度
	循环整个数组，比较当前元素和指定val，不相等的话，是有效的元素，将其移动到pos位置
*/
package questionBank

import (
	"math/rand"
	"time"
)

type Solution struct {
	origin []int
	result []int
}

func Constructor(nums []int) Solution {
	result := make([]int, len(nums))
	//这里需要复制一份结果出来，因为切片使用同一份数据，对一个切片的修改会显示到另外的切片上面，影响数据
	copy(result, nums)
	solution := Solution{
		origin : nums,
		result : result,
	}

	return solution
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	this.result = this.origin

	return this.result
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	copy(this.result, this.origin)
	rand.Seed(time.Now().Unix())
	for index := range this.origin {
		change := rand.Intn(len(this.origin))
		this.result[index], this.result[change] = this.result[change], this.result[index]
	}

	return this.result
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func ShuffleAnArray(nums []int) ([]int, []int) {
	obj := Constructor(nums)
	param_1 := obj.Shuffle()
	param_2 := obj.Reset()

	return param_1, param_2
}