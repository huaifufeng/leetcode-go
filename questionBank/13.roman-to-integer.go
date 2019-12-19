/**
  题目地址：https://leetcode-cn.com/problems/roman-to-integer/

  罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

  <pre>
     字符          数值
     I             1
     V             5
     X             10
     L             50
     C             100
     D             500
     M             1000
  </pre>

  <pre>
     例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

     通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
         I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
         X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
         C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
  </pre>

  给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

  解题： 时间复杂度O(n)
    建立罗马数字和整数的对应关系表，然后从左向右依次处理每个罗马数字，如果遇到I X C 就在判断下一个罗马数字是不是相对应 VX LC DM, 如果是就是进行减法，否则就是进行加法
*/
package questionBank

func romanToInt(s string) int {
	strLen := len(s)
 	intNum := 0
 	mapStr := map[byte]int{
 		'I' : 1,
 		'V' : 5,
 		'X' : 10,
 		'L' : 50,
 		'C' : 100,
 		'D' : 500,
 		'M' : 1000,
	}

	for i := 0; i < strLen; i++ {
		if i + 1 == strLen {
			intNum += mapStr[s[i]]
			break
		} else {
			if (s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X')) || (s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C')) || (s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M')) {
				intNum -= mapStr[s[i]]
			} else {
				intNum += mapStr[s[i]]
			}
		}
	}

	return intNum
}

func RomanToInt(s string) int {
	return romanToInt(s)
}