package binary_search_tree

import "github.com/william-lg/go-exercises/data_structure/utils"

func inorderWalk(root *TreeNode) (nodes []int) {
	if root == nil {
		return
	}

	nodes = append(nodes, inorderWalk(root.Left)...)
	nodes = append(nodes, root.Key)
	nodes = append(nodes, inorderWalk(root.Right)...)

	return
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + utils.MaxInt(height(root.Left), height(root.Right))
}
