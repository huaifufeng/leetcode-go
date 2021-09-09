/**
  题目地址：https://leetcode-cn.com/problems/text-justification/

  解题：
  1、循环字符串数组，找到每个符合条件的字符串子数组，拼接字符串。O(n^2) 循环整个数组，当子字符串数组就是整个数组时，内部也是n的循环
*/
package questionBank

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	//有效单词的长度
	wordLength := 0
	//有效单词在字符串数组的左右索引值
	l := 0
	r := 0
	res := make([]string, 0)
	for i, word := range words {
		//当选中一个以上单词之后，每次计算新单词长度时，需要加上一个单词间隔空格，
		//相等的时候表示第一个单词，第一个单词不需要
		newNum := 0
		if i > 0 {
			newNum = 1
		}
		//计算放下这么多单词的长度，如果大于最大长度，说明放不下了，已经是一个有效的字符串了
		if len(word)+wordLength+r-l+newNum > maxWidth {
			//获取有效字符串数组，需要额外的空白数量
			subWords := words[l : r+1]
			SpaceLength := maxWidth - wordLength - (r - l)
			res = append(res, getString(subWords, SpaceLength, false))

			//初始化有效字符串在字符串数组中的索引和有效单词长度
			l = r + 1
			r = l
			wordLength = len(word)
		} else {
			//不大于最大长度时，可以单词数组索引右移，单词长度增加
			wordLength += len(word)
			r = i
		}
	}

	//单词数组完毕之后，进行最后余下字符串的处理
	subWords := words[l:]
	SpaceLength := maxWidth - wordLength - (r - l)
	res = append(res, getString(subWords, SpaceLength, true))

	return res
}

//getString  获取字符串子数组拼接的字符串
func getString(subWords []string, SpaceLength int, isLast bool) string {
	//只有一个单词的话，左对齐处理
	if len(subWords) == 1 {
		return subWords[0] + strings.Repeat(" ", SpaceLength)
	}

	//最后一个单词的话，左对齐处理
	if isLast {
		return strings.Join(subWords, " ") + strings.Repeat(" ", SpaceLength)
	}

	//获取单词间平均放置的空格，和额外多出的空格数量
	num := SpaceLength / (len(subWords) - 1)
	left := SpaceLength % (len(subWords) - 1)
	str := ""
	//循环前n-1个单词依次拼接相应的平均空格数量，对于多出的空格，依次向后摆放
	for i := 0; i < len(subWords)-1; i++ {
		str = str + subWords[i] + strings.Repeat(" ", num+1)
		if left > 0 {
			str += " "
			left--
		}
	}

	//最后一个单词单独存放，因为需要右对齐
	str += subWords[len(subWords)-1]
	return str
}
