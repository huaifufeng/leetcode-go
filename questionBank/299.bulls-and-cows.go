/**
  题目地址：https://leetcode-cn.com/problems/bulls-and-cows/

  解题：
  1、
*/
package questionBank

import (
	"fmt"
	"strconv"
)

func getHint(secret, guess string) string {
	bulls := 0
	var cntS, cntG [10]int
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			cntS[secret[i]-'0']++
			cntG[guess[i]-'0']++
		}
	}
	cows := 0
	for i := 0; i < 10; i++ {
		cows += min(cntS[i], cntG[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getHint1(secret string, guess string) string {
	sMaps := make(map[byte]map[int]int)
	gMaps := make(map[byte][]int)
	for i := 0; i < len(secret); i++ {
		if len(sMaps[secret[i]]) == 0 {
			sMaps[secret[i]] = make(map[int]int)
		}
		sMaps[secret[i]][i] = 1
		gMaps[guess[i]] = append(gMaps[guess[i]], i)
	}

	a := 0
	b := 0
	for v, iSlice := range gMaps {
		sMap := sMaps[v]
		ai := 0
		for _, iv := range iSlice {
			if sMap[iv] == 1 {
				ai++
			}
		}

		length := 0
		if len(iSlice) > len(sMap) {
			length = len(sMap)
		} else {
			length = len(iSlice)
		}

		b += length - ai
		a += ai
	}

	return fmt.Sprintf("%sA%sB", strconv.Itoa(a), strconv.Itoa(b))
}
