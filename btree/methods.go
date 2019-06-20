package btree

func (tree *BTree) isFull(node *Node) bool {
	return len(node.Keys) >= tree.maxKeysCount()
}

func (tree *BTree) maxKeysCount() int {
	return int(tree.Degree) - 1
}

func (tree *BTree) minChildrenCount() int {
	return int((tree.Degree + 1) / 2)
}

func (tree *BTree) minKeysCount() int {
	return tree.minChildrenCount() - 1
}

func (node *Node) IsLeaf() bool {
	return len(node.Children) == 0
}
