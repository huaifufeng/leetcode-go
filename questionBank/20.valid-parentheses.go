/**
  题目地址：https://leetcode-cn.com/problems/valid-parentheses/

  给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

  <pre>
     有效字符串需满足：
       1.左括号必须用相同类型的右括号闭合。
       2.左括号必须以正确的顺序闭合。

     注意空字符串可被认为是有效字符串。
  </pre>

  <pre>
     输入: "()[]{}"
     输出: true

     输入: "{[]}"
     输出: true
  </pre>


  解题： 时间复杂度O(n)
	依次获取字符串中的每个字符，将每个不是结束的字符放到堆栈中，如果是结束的字符的话，就判断一下最后一个字符和这个字符能不能匹配，
    可以匹配就把最后一个字符从堆栈中移除，否则就是失败了。最后，如果堆栈中没有字符了，说明匹配正确。
*/
package questionBank

func isValid(s string) bool {
	length := len(s)
	//空字符返回true
	if length == 0 {
		return true
	}

	//奇数说明没有全部配对，返回false
	if length%2 == 1 {
		return false
	}

	strMap := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	strSlice := make([]byte, length)
	left := 0
	for i := 0; i < length; i++ {
		oneB, ok := strMap[s[i]]
		if ok {
			if left == 0 || strSlice[left-1] != oneB {
				return false
			} else {
				left--
			}
		} else {
			strSlice[left] = s[i]
			left++
		}
	}

	if left != 0 {
		return false
	}

	return true
}
