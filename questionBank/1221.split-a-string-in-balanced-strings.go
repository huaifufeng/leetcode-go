/**
  题目地址：https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/

  在一个 平衡字符串 中，'L' 和 'R' 字符的数量是相同的。
  给你一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。
  注意：分割得到的每个字符串都必须是平衡字符串。
  返回可以通过分割得到的平衡字符串的 最大数量 。

  <pre>
     输入：s = "RLRRLLRLRL"
     输出：4
     解释：s 可以分割为 "RL"、"RRLL"、"RL"、"RL" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。
  </pre>

  <pre>
     输入：s = "RLLLLRRRLR"
     输出：3
     解释：s 可以分割为 "RL"、"LLLRRR"、"LR" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。
  </pre>

  <pre>
     输入：s = "LLLLRRRR"
     输出：1
     解释：s 只能保持原样 "LLLLRRRR".
  </pre>

  <pre>
     输入：s = "RLRRRLLRLL"
     输出：2
     解释：s 可以分割为 "RL"、"RRRLLRLL" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。
  </pre>

  PS:
    1 <= s.length <= 1000
    s[i] = 'L' 或 'R'
    s 是一个 平衡 字符串

  解题：
	1、因为s是平衡字符串，说明具备相同数量的R、L，从左向右循环，每次匹配到平衡字符串之后计数，并重新开始匹配，O(n)
*/
package questionBank

func balancedStringSplit(s string) int {
	total := 0
	rnum := 0
	lnum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			rnum++
		}

		if s[i] == 'L' {
			lnum++
		}

		if rnum > 0 && rnum == lnum {
			total++
			rnum = 0
			lnum = 0
		}
	}

	return total
}
