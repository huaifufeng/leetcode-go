/**
  题目地址：https://leetcode-cn.com/problems/smallest-k-lcci/

  设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。

  <pre>
     输入： arr = [1,3,5,7,2,4,6,8], k = 4
     输出： [1,2,3,4]
  </pre>

  PS:
    0 <= len(arr) <= 100000
    0 <= k <= min(100000, len(arr))

  解题：
    1.用go本身的方法，先排序，在取k个元素
    2.
*/
package interviewQuestion

import "sort"

func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
