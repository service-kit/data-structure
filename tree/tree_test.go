package tree

import (
	"testing"
	"fmt"
)

func Test_BinaryTree(t *testing.T) {
	root := CreateBinaryTreeNode(1)
	f1 := CreateBinaryTreeNode(2)
	root.SetLeft(f1)
	f2 := CreateBinaryTreeNode(3)
	root.SetRight(f2)
	s1 := CreateBinaryTreeNode(4)
	s2 := CreateBinaryTreeNode(5)
	f1.SetLeft(s1)
	f1.SetRight(s2)
	s3 := CreateBinaryTreeNode(6)
	s4 := CreateBinaryTreeNode(7)
	f2.SetLeft(s3)
	f2.SetRight(s4)
	fmt.Println("pre order")
	root.PreOrder()
	fmt.Println("\nmiddle order")
	root.MiddleOrder()
	fmt.Println("\npost order")
	root.PostOrder()
	fmt.Println()
}