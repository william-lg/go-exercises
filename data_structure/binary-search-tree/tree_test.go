package binary_search_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_TreeInsert(t *testing.T) {
	tree := new(Tree)
	var node *TreeNode

	node = &TreeNode{Key: 4}
	tree.TreeInsert(node)
	printArrayInt(tree.InorderWalk())

	node = &TreeNode{Key: 1}
	tree.TreeInsert(node)
	printArrayInt(tree.InorderWalk())

	node = &TreeNode{Key: 2}
	tree.TreeInsert(node)
	printArrayInt(tree.InorderWalk())

	node = &TreeNode{Key: 6}
	tree.TreeInsert(node)
	printArrayInt(tree.InorderWalk())

	node = &TreeNode{Key: 7}
	tree.TreeInsert(node)
	printArrayInt(tree.InorderWalk())

	node = &TreeNode{Key: 5}
	tree.TreeInsert(node)
	printArrayInt(tree.InorderWalk())

	node = &TreeNode{Key: 4}
	tree.TreeInsert(node)
	printArrayInt(tree.InorderWalk())
}

func printArrayInt(arr []int) {
	for _, v := range arr {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func TestTree_TreeDelete(t *testing.T) {
	tree := new(Tree)
	var node *TreeNode

	node = &TreeNode{Key: 4}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 2}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 1}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 3}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 6}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 5}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 10}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 8}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 9}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 11}
	tree.TreeInsert(node)

	var tar *TreeNode
	tar = tree.TreeSearch(1)
	tree.TreeDelete(tar)
	printArrayInt(tree.InorderWalk())

	tar = tree.TreeSearch(2)
	tree.TreeDelete(tar)
	printArrayInt(tree.InorderWalk())

	tar = tree.TreeSearch(6)
	tree.TreeDelete(tar)
	printArrayInt(tree.InorderWalk())

	tar = tree.TreeSearch(8)
	tree.TreeDelete(tar)
	printArrayInt(tree.InorderWalk())
}

func TestTree_TreeSearch(t *testing.T) {
	tree := new(Tree)
	var node *TreeNode

	node = &TreeNode{Key: 4}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 2}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 1}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 3}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 6}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 5}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 10}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 8}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 9}
	tree.TreeInsert(node)

	var tar *TreeNode
	tar = tree.TreeSearch(4)
	assert.Equal(t, 4, tar.Key)
	tar = tree.TreeSearch(9)
	assert.Equal(t, 9, tar.Key)
	tar = tree.TreeSearch(6)
	assert.Equal(t, 6, tar.Key)
}

func TestTree_TreeSuccessor(t *testing.T) {
	tree := new(Tree)
	var node *TreeNode

	node = &TreeNode{Key: 4}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 2}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 1}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 3}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 6}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 5}
	tree.TreeInsert(node)

	var tar, sucor *TreeNode
	tar = tree.TreeSearch(4)
	sucor = tree.TreeSuccessor(tar)
	assert.Equal(t, 5, sucor.Key)
	tar = tree.TreeSearch(3)
	sucor = tree.TreeSuccessor(tar)
	assert.Equal(t, 4, sucor.Key)
}

func TestTree_Height(t *testing.T) {
	tree := new(Tree)
	var node *TreeNode

	node = &TreeNode{Key: 4}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 2}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 1}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 3}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 6}
	tree.TreeInsert(node)

	height := tree.Height()
	assert.Equal(t, 3, height)

	node = &TreeNode{Key: 7}
	tree.TreeInsert(node)
	node = &TreeNode{Key: 8}
	tree.TreeInsert(node)
	height = tree.Height()
	assert.Equal(t, 4, height)
}
