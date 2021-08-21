/**
  题目地址：https://leetcode-cn.com/problems/string-compression/

  给你一个字符数组 chars ，请使用下述算法压缩：
    从一个空字符串 s 开始。对于 chars 中的每组 连续重复字符 ：
      如果这一组长度为 1 ，则将字符追加到 s 中。
      否则，需要向 s 追加字符，后跟这一组的长度。
  压缩后得到的字符串 s 不应该直接返回 ，需要转储到字符数组 chars 中。需要注意的是，如果组长度为 10 或 10 以上，则在 chars 数组中会被拆分为多个字符。

  请在 修改完输入数组后 ，返回该数组的新长度。
  你必须设计并实现一个只使用常量额外空间的算法来解决此问题。

  <pre>
    输入：chars = ["a","a","b","b","c","c","c"]
    输出：返回 6 ，输入数组的前 6 个字符应该是：["a","2","b","2","c","3"]
    解释：
        "aa" 被 "a2" 替代。"bb" 被 "b2" 替代。"ccc" 被 "c3" 替代。
  </pre>

  <pre>
    输入：chars = ["a"]
    输出：返回 1 ，输入数组的前 1 个字符应该是：["a"]
    解释：
        没有任何字符串被替代。
  </pre>

  <pre>
    输入：chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
    输出：返回 4 ，输入数组的前 4 个字符应该是：["a","b","1","2"]。
    解释：
        由于字符 "a" 不重复，所以不会被压缩。"bbbbbbbbbbbb" 被 “b12” 替代。注意每个数字在数组中都有它自己的位置。
  </pre>

  PS：
    1 <= chars.length <= 2000
    chars[i] 可以是小写英文字母、大写英文字母、数字或符号

  解题：
    1、循环整个字符串数组，依次比较当前元素和之前元素的值，进行元素值得计数。O(n)
*/
package questionBank

import "strconv"

func compress(chars []byte) int {
	var lastChar byte
	index := 0
	num := 0
	for _, char := range chars {
		if lastChar != char {
			if num > 0 {
				chars[index] = lastChar
				index++
				if num > 1 {
					numStr := strconv.Itoa(num)
					copy(chars[index:index+len(numStr)], numStr)
					index += len(numStr)
				}
			}

			lastChar = char
			num = 1

			continue
		}

		num++
	}

	if num > 0 {
		chars[index] = lastChar
		index++
		if num > 1 {
			numStr := strconv.Itoa(num)
			copy(chars[index:index+len(numStr)], numStr)
			index += len(numStr)
		}
	}

	chars = chars[:index]

	return len(chars)
}
