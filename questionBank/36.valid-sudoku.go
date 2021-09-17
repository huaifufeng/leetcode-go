/**
  题目地址：https://leetcode-cn.com/problems/valid-sudoku/


  解题：
	1、变量二维数组，分别建立行，列，方块的元素映射，只要存在大于1的值，说明重复非法， O(1) = O(81)
*/
package questionBank

func isValidSudoku(board [][]byte) bool {
	//方块
	mr := make(map[uint8]map[byte]uint8)
	//列
	mc := make(map[uint8]map[byte]uint8)
	for i, i2 := range board {
		//行
		ml := make(map[byte]uint8)
		for i3, b := range i2 {
			if b == '.' {
				continue
			}
			ml[b]++
			if ml[b] > 1 {
				return false
			}

			mkey := uint8(i/3*10 + i3/3)
			if _, ok := mr[mkey]; !ok {
				mr[mkey] = make(map[byte]uint8)
			}
			mr[mkey][b]++
			if mr[mkey][b] > 1 {
				return false
			}

			if _, ok := mc[uint8(i3)]; !ok {
				mc[uint8(i3)] = make(map[byte]uint8)
			}
			mc[uint8(i3)][b]++
			if mc[uint8(i3)][b] > 1 {
				return false
			}
		}
	}

	return true
}
