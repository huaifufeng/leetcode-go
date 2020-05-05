package questionBank

import "testing"

func TestIsValidBSTNot(t *testing.T) {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
		},
	}

	res := isValidBST(tree)
	if res {
		t.Error("判断是否为有效二叉搜索树的方法错误！")
	}
}

func TestIsValidBSTYes(t *testing.T) {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	res := isValidBST(tree)
	if !res {
		t.Error("判断是否为有效二叉搜索树的方法错误！")
	}
}
