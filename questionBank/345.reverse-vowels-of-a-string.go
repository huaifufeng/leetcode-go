/**
  题目地址：https://leetcode-cn.com/problems/reverse-vowels-of-a-string/

  反转字符串中的元音字母。
  编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

  <pre>
    输入："hello"
    输出："holle"
  </pre>

  解题： 时间复杂度O(n2) n为字符串长度
	从左右两个方向同时开始循环，比较找到的字符类型，如果都是元音字符进行交互，这里注意大写字母
*/
package questionBank

func reverseVowels(s string) string {
	vowels := map[byte]int8{
		'a': 1,
		'A': 1,
		'e': 1,
		'E': 1,
		'i': 1,
		'I': 1,
		'o': 1,
		'O': 1,
		'u': 1,
		'U': 1,
	}
	if len(s) == 0 {
		return s
	}

	right := len(s) - 1
	str := []byte(s)
	for i := 0; i < len(s); i++ {
		if i >= right {
			break
		}
		if _, ok := vowels[s[i]]; ok {
			leftVal := s[i]
			for j := right; j > 0; j-- {
				if i == j {
					right = j
					break
				}
				if _, ok := vowels[s[j]]; ok {
					right = j - 1
					str[i] = s[j]
					str[j] = leftVal
					break
				}
			}
		}
	}

	return string(str)
}
