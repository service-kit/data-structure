package tree

type BinarySearchTree struct {
	root *BinaryTreeNode
}

func (b *BinarySearchTree) Add(data int) {
	if nil == b.root {
		b.root = CreateBinaryTreeNode(data)
		return
	}
	pos := b.root
	for {
		if pos.value > data {
			pos = pos.left
		}else {
			pos = pos.right
		}
	}
}