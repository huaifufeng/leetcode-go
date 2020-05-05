/**
  给定一个二叉树，判断其是否是一个有效的二叉搜索树。
  假设一个二叉搜索树具有如下特征：
    节点的左子树只包含小于当前节点的数。
    节点的右子树只包含大于当前节点的数。
    所有左子树和右子树自身必须也是二叉搜索树。

  <pre>
     输入:
      2
     / \
    1   3
    输出: true
  </pre>

  解法：
  1、依次处理树的每一个节点，比较节点的值是否满足要求，这里需要注意的是：左节点要小于父节点，右节点要大于父节点的值，这里
  指的是包括在父节点的父节点的值在内的值，所以左节点的最大值和右节点的最小值都要不断变化，知道满足条件
*/
package questionBank

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return subValid(root, math.MinInt64, math.MaxInt64)
}

func subValid(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}

	if node.Val <= lower || node.Val >= upper {
		return false
	}

	return subValid(node.Left, lower, node.Val) && subValid(node.Right, node.Val, upper)
}
