package tree

import "fmt"

type BinaryTreeNode struct {
	pre *BinaryTreeNode
	left *BinaryTreeNode
	right *BinaryTreeNode
	value int
}

func (b *BinaryTreeNode) Pre() *BinaryTreeNode {
	return b.pre
}

func (b *BinaryTreeNode) Left() *BinaryTreeNode {
	return b.left
}

func (b *BinaryTreeNode) Right() *BinaryTreeNode {
	return b.right
}

func (b *BinaryTreeNode) Value() int {
	return b.value
}

func (b *BinaryTreeNode) SetLeft(l *BinaryTreeNode) {
	b.left = l
	l.pre = b
}

func (b *BinaryTreeNode) SetRight(r *BinaryTreeNode) {
	b.right = r
	r.pre = b
}

func (b *BinaryTreeNode) PreOrder() {
	if nil == b {
		return
	}
	fmt.Print(b.value)
	b.left.PreOrder()
	b.right.PreOrder()
}

func (b *BinaryTreeNode) MiddleOrder() {
	if nil == b {
		return
	}
	b.left.MiddleOrder()
	fmt.Print(b.value)
	b.right.MiddleOrder()
}

func (b *BinaryTreeNode) PostOrder() {
	if nil == b {
		return
	}
	b.left.PostOrder()
	b.right.PostOrder()
	fmt.Print(b.value)
}

func CreateBinaryTreeNode(data int) *BinaryTreeNode {
	b := new(BinaryTreeNode)
	b.value = data
	return b
}