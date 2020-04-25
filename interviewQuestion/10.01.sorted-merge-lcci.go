/**
  题目地址：https://leetcode-cn.com/problems/sorted-merge-lcci/
*/
package interviewQuestion

func merge(A []int, m int, B []int, n int) {
	pos := m + n - 1

	for m > 0 || n > 0 {
		if m > 0 && n > 0 {
			if A[m-1] >= B[n-1] {
				A[pos] = A[m-1]
				m--
			} else {
				A[pos] = B[n-1]
				n--
			}
		} else if m > 0 && n == 0 {
			A[pos] = A[m-1]
			m--
		} else {
			A[pos] = B[n-1]
			n--
		}

		pos--
	}
}

func Merge(A []int, m int, B []int, n int) []int {
	merge(A, m, B, n)

	return A
}
