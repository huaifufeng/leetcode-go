/**
 题目地址：https://leetcode-cn.com/problems/happy-number/

 编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，
 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，
 那么这个数就是快乐数。
 如果 n 是快乐数就返回 True ；不是，则返回 False 。

 <pre>
 输入：19
 输出：true
 解释：
   12 + 92 = 82
   82 + 22 = 68
   62 + 82 = 100
   12 + 02 + 02 = 1
 </pre>

解题： 时间复杂度O(n) n为数组长度

*/
package questionBank

//用map来存放已经有的数值，如果遇到map中已有的值，说明遇到了循环，不用再进行处理了
func isHappy(n int) bool {
	hasMap := make(map[int]int)
	for n > 0 {
		n = getSquareSum(n)

		if n == 1 {
			return true
		}

		if _, ok := hasMap[n]; ok {
			break
		}

		hasMap[n] = 1
	}

	return false
}

func getSquareSum(num int) int {
	sum := 0
	for num > 0 {
		cur := num % 10
		sum += cur * cur
		num = num / 10
	}

	return sum
}

//快慢指针法，类似于判断链表中是否有环的方法
func isHappy2(n int) bool {
	slow := n
	fast := getSquareSum(n)

	//当快指针等于1 说明获得了想要的值
	//当快指针等于慢指针时，说明遇到了环，已经失败
	for fast != 1 && fast != slow {
		slow = getSquareSum(slow)
		fast = getSquareSum(getSquareSum(fast))
	}

	return fast == 1
}

//4, 16, 37, 58, 89, 145, 42, 20  循环链的值，在这些值中将会失败
