package binary_search_tree

type TreeNode struct {
	Key    int
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

func (tn *TreeNode) IsLeaf() bool {
	if tn == nil {
		return false
	}

	return tn.Left == nil && tn.Right == nil
}
