package binary_search_tree

type Tree struct {
	Root *TreeNode
}

func (tree *Tree) IsEmpty() bool {
	return tree.Root == nil
}

func (tree *Tree) Height() int {
	return height(tree.Root)
}

func (tree *Tree) InorderWalk() (nodes []int) {
	return inorderWalk(tree.Root)
}

func (tree *Tree) TreeInsert(node *TreeNode) {
	var it, itp *TreeNode

	for it = tree.Root; it != nil; {
		itp = it
		if node.Key <= it.Key {
			it = it.Left
		} else {
			it = it.Right
		}
	}

	node.Parent = itp
	if itp == nil {
		tree.Root = node
	} else if node.Key <= itp.Key {
		itp.Left = node
	} else {
		itp.Right = node
	}

	return
}

func (tree *Tree) TreeTransplant(old, new1 *TreeNode) {
	if tree.IsEmpty() {
		tree.Root = new1
		return
	}

	parent := old.Parent
	if old == parent.Left {
		parent.Left = new1
	} else if old == parent.Right {
		parent.Right = new1
	}
	if new1 != nil {
		new1.Parent = parent
	}
}

func (tree *Tree) TreeDelete(tar *TreeNode) {
	if tar == nil {
		return
	}

	if tar.Left == nil {
		tree.TreeTransplant(tar, tar.Right)
	} else if tar.Right == nil {
		tree.TreeTransplant(tar, tar.Left)
	} else {
		sucor := tree.TreeMinimum(tar.Right)
		if sucor != tar.Right {
			tree.TreeTransplant(sucor, sucor.Right)
			sucor.Right = tar.Right
			sucor.Right.Parent = sucor
		}
		tree.TreeTransplant(tar, sucor)
		sucor.Left = tar.Left
		sucor.Left.Parent = sucor
	}
}

func (tree *Tree) TreeMinimum(x *TreeNode) *TreeNode {
	for ; x.Left != nil; x = x.Left {
	}

	return x
}

func (tree *Tree) TreeSearch(key int) *TreeNode {
	it := tree.Root
	for it != nil && it.Key != key {
		if key <= it.Key {
			it = it.Left
		} else {
			it = it.Right
		}
	}

	return it
}

func (tree *Tree) TreeSuccessor(tar *TreeNode) *TreeNode {
	if tar.Right != nil {
		return tree.TreeMinimum(tar.Right)
	}

	y := tar.Parent
	for y != nil && tar == y.Right {
		y = y.Parent
		tar = tar.Parent
	}

	return y
}
