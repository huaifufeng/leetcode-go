//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/31/
package 数组

//因为已经有这个方法名称的，这里末尾加了个0
//对换每行，然后对换对角线
func rotate0(matrix [][]int) {
	length := len(matrix)

	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-1-i] = matrix[length-1-i], matrix[i]
	}

	for i := 0; i < length; i++ {
		for k := i; k < length; k++ {
			matrix[i][k], matrix[k][i] = matrix[k][i], matrix[i][k]
		}
	}
}

//因为已经有这个方法名称的，这里末尾加了个0
//左右对换列，然后对换对角线
func rotate01(matrix [][]int) {
	length := len(matrix)
	//翻转每一行的数据
	for i := 0; i < length; i++ {
		for j := length - 1; j > length/2-1; j-- {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}

	for i := length - 2; i >= 0; i-- {
		for j := 0; j < length-1-i; j++ {
			matrix[i][j], matrix[length-1-j][length-1-i] = matrix[length-1-j][length-1-i], matrix[i][j]
		}
	}
}
