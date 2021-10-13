/**
  题目地址：https://leetcode-cn.com/problems/fizz-buzz/

  解题：
    1、循环1~n，判断每个数字和3，5的关系，防入数组 O(n)
*/
package questionBank

import "strconv"

func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 1; i <= n; i++ {
		left3 := i % 3
		left5 := i % 5
		if left3 == 0 {
			if left5 == 0 {
				answer[i-1] = "FizzBuzz"
			} else {
				answer[i-1] = "Fizz"
			}
		} else {
			if left5 == 0 {
				answer[i-1] = "Buzz"
			} else {
				answer[i-1] = strconv.Itoa(i)
			}
		}
	}

	return answer
}
