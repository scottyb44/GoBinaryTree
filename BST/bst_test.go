package bst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func DefaultSetup() *BST {
	bst := NewBST(5)

	bst.InsertFromRoot(4)
	bst.InsertFromRoot(8)
	bst.InsertFromRoot(9)

	return bst
}

// TODO SB: This testing isnt good
func TestBST_Insert_NodeExist(t *testing.T) {
	bst := NewBST(5)

	bst.InsertFromRoot(4)
	bst.InsertFromRoot(8)
	bst.InsertFromRoot(9)

	require.NotNil(t, bst.root.leftChild)
	require.NotNil(t, bst.root.rightChild)
	require.Nil(t, bst.root.leftChild.leftChild)
	require.Nil(t, bst.root.leftChild.rightChild)
	require.Nil(t, bst.root.rightChild.leftChild)
	require.NotNil(t, bst.root.rightChild.rightChild)
}

func TestBST_Insert_HasCorrectParent(t *testing.T) {
	bst := NewBST(5)

	node := bst.InsertFromRoot(4)
	require.Equal(t, 5, node.parent.value)

	node = bst.InsertFromRoot(8)
	require.Equal(t, 5, node.parent.value)

	node = bst.InsertFromRoot(9)
	require.Equal(t, 8, node.parent.value)
}

func TestBST_FindNode_NotExistent(t *testing.T) {
	bst := DefaultSetup()

	node := bst.FindNode(10)

	require.Nil(t, node)
}

func TestBST_FindNode_Existent(t *testing.T) {
	bst := DefaultSetup()

	node := bst.FindNode(9)

	require.NotNil(t, node)
	require.Equal(t, 9, node.value)
}