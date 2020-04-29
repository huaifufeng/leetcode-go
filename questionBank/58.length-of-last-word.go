/**
  题目地址：https://leetcode-cn.com/problems/length-of-last-word/

  给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
  如果不存在最后一个单词，请返回 0 。
  说明：一个单词是指由字母组成，但不包含任何空格的字符串。

  <pre>
    输入: "Hello World"
    输出: 5
  </pre>

  解题：
    1、处理特殊情况之后，直接使用strings带的方法将字符串转换成单词的切片，直接获取最后一个单词的长度
	2、遍历整个字符串，遇到空字符串，说明接下来的遇到字符串是需要修改单词的长度的，这里就设置标识为需要修改；
       否则，就按照标识如果需要修改，就设置单词长度为1，并标识设置为0不需要修改，直到遇到空字符串；否则就是增加单词长度

*/
package questionBank

import (
	"strings"
)

//使用go原生的字符串分隔方法
func lengthOfLastWord(s string) int {
	strLength := len(s)
	//空字符串直接返回0
	if strLength == 0 {
		return 0
	}

	//将字符串转换为单词切片，处理为空的情况
	words := strings.Fields(s)
	if len(words) == 0 {
		return 0
	}

	return len(words[len(words)-1])
}

func lengthOfLastWord2(s string) int {
	strLength := len(s)

	needChange := 1
	curLength := 0
	for i := 0; i < strLength; i++ {
		//遇到空字符串，说明接下来的遇到字符串是需要修改单词的长度的，这里就设置标识为需要修改
		if s[i] == ' ' {
			if needChange != 1 {
				needChange = 1
			}
			//否则，就按照标识如果需要修改，就设置单词长度为1，并标识设置为0不需要修改，直到遇到空字符串；否则就是增加单词长度
		} else {
			if needChange == 1 {
				curLength = 1
				needChange = 0
			} else {
				curLength++
			}
		}
	}

	return curLength
}
