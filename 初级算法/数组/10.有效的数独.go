//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/30/
package 数组

func isValidSudoku(board [][]byte) bool {
	verticalMap := make([]map[byte]int, 9)
	rectMap := make([]map[byte]int, 3)

	for i:=0; i<9; i++ {
		lineMap := make(map[byte]int, 9)

		if i % 3 == 0 {
			rectMap = make([]map[byte]int, 3)
		}

		for j:= 0; j<9;j++ {
			if board[i][j] == '.' {
				continue
			}

			if _, ok := lineMap[board[i][j]];ok {
				return false
			} else {
				lineMap[board[i][j]] = 1
			}

			if _, ok := verticalMap[j][board[i][j]]; ok {
				return false
			} else {
				if verticalMap[j] == nil {
					verticalMap[j] = make(map[byte]int)
				}
				verticalMap[j][board[i][j]] = 1
			}

			key := 0
			if j >= 3 && j <6 {
				key = 1
			} else if j>=6 {
				key =2
			}

			if _, ok := rectMap[key][board[i][j]];ok {
				return false
			} else {
				if rectMap[key] == nil {
					rectMap[key] = make(map[byte]int)
				}
				rectMap[key][board[i][j]] = 1
			}
		}
	}

	return true
}
