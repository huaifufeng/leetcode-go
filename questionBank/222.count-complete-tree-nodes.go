/**
  题目地址：https://leetcode-cn.com/problems/count-complete-tree-nodes/

  给出一个完全二叉树，求出该树的节点个数。

  说明：
    完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
    并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~2h个节点。

  <pre>
    输入:
      1
     / \
    2   3
   / \  /
  4  5 6

  输出: 6
  </pre>

  解题：使用递归操作，分别求当前节点的左边节点数和右边节点数，然后想加到一起。
*/
package questionBank

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	count := 0
	if root == nil {
		return count
	}

	var left *TreeNode = root.Left
	var right *TreeNode = root.Right
	count++

	count += countNodes(left)
	count += countNodes(right)

	return count
}
