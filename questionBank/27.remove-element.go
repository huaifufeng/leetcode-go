/**
  题目地址：https://leetcode-cn.com/problems/remove-element/

  给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
  不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
  元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

  <pre>
    给定 nums = [0,1,2,2,3,0,4,2], val = 2,
    函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
    注意这五个元素可为任意顺序。
    你不需要考虑数组中超出新长度后面的元素。
  </pre>

  解题： 时间复杂度O(n) n为数组长度
	循环整个数组，比较当前元素和指定val，不相等的话，是有效的元素，将其移动到pos位置
*/
package questionBank

func removeElement(nums []int, val int) int {
	pos := 0

	for i := 0; i < len(nums); i++ {
		//如果当前元素不等于指定的值，说明这个元素是需要替换的元素，将元素替换到pos的位置
		if nums[i] != val {
			nums[pos] = nums[i]
			pos++
		}
	}

	return pos
}

func RemoveElement(nums []int, val int) int {
	return removeElement(nums, val)
}
