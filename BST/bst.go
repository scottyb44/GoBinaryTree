package bst

type TreeNode struct {
	value      int
	parent     *TreeNode
	leftChild  *TreeNode
	rightChild *TreeNode
}

func NewTreeNode(val int, par *TreeNode) *TreeNode {
	treeNode := TreeNode{value: val, parent: par, leftChild: nil, rightChild: nil}
	return &treeNode
}

// TODO SB: This is probably not a useful abstraction
type BST struct {
	root TreeNode
}

func NewBST(rootVal int) *BST {
	rootNode := NewTreeNode(rootVal, nil)
	bst := BST{root: *rootNode}
	return &bst
}

// TODO SB: Probably dont need to return the bst
func (b *BST) InsertFromRoot(val int) *TreeNode {
	return Insert(val, &b.root)
}

// TODO SB: What about a find from a given node, will need to use parent field
func (b *BST) FindNode(val int) *TreeNode {
	return FindFromGivenNode(val, &b.root)
}

func FindFromGivenNode(val int, current *TreeNode) *TreeNode {
	if nil == current {
		return nil
	}
	if val == current.value {
		return current
	}
	if val < current.value {
		return FindFromGivenNode(val, current.leftChild)
	} else {
		return FindFromGivenNode(val, current.rightChild)
	}
}

// TODO SB: This function will need some cleanup
func Insert(val int, parent *TreeNode) *TreeNode {
	if val <= parent.value {
		if parent.leftChild == nil {
			newNode := NewTreeNode(val, parent)
			parent.leftChild = newNode
			return newNode
		}
		return Insert(val, parent.leftChild)
	} else {
		if parent.rightChild == nil {
			newNode := NewTreeNode(val, parent)
			parent.rightChild = newNode
			return newNode
		}
		return Insert(val, parent.rightChild)
	}
}
