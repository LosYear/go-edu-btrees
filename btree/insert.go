package btree

import "log"

func (tree *BTree) Insert(key int) {
	leafNode, found := searchRoutine(tree.Root, key)

	if found {
		log.Fatalln("Key already exists in tree")
	}

	newHead := tree.insertionRoutine(leafNode, key)

	if newHead != nil {
		tree.Root = newHead
	}
}

func (tree *BTree) insertionRoutine(node *Node, key int) *Node {
	// If empty tree
	if node == nil {
		return &Node{Keys: []int{key}}
	}

	// If node is full
	if !tree.isFull(node) {
		node.insertKey(key)
	} else {
		node.insertKey(key)
		median, left, right := tree.splitNode(node)

		// Replace node by two in child list
		if node.Parent != nil {
			replaceNode(node, left, right)
			return tree.insertionRoutine(node.Parent, median)

		} else {
			// New Root
			newNode := &Node{Keys: []int{median}, Children: []*Node{left, right}}
			fixParents(newNode, &newNode.Children)

			return newNode
		}
	}

	return nil
}

func (tree *BTree) splitNode(node *Node) (median int, left, right *Node) {
	medianIndex := len(node.Keys) / 2
	median = node.Keys[medianIndex]

	left = &Node{Keys: make([]int, 0, tree.maxKeysCount()), Children: make([]*Node, 0, tree.maxKeysCount()+1)}

	left.Keys = append(left.Keys, node.Keys[:medianIndex]...)
	if !node.IsLeaf() {
		left.Children = append(left.Children, node.Children[:len(left.Keys)+1]...)
		fixParents(left, &left.Children)
	}

	right = &Node{Keys: make([]int, 0, tree.maxKeysCount()), Children: make([]*Node, 0, tree.maxKeysCount()+1)}
	right.Keys = append(right.Keys, node.Keys[medianIndex+1:]...)

	if !node.IsLeaf() {
		right.Children = append(right.Children, node.Children[len(left.Keys)+1:]...)
		fixParents(right, &right.Children)
	}

	return median, left, right
}

func replaceNode(node, left, right *Node) {
	parent := node.Parent
	index := node.nodeIndex()

	parent.Children = append(parent.Children, nil)
	copy(parent.Children[index+1:], parent.Children[index:])

	parent.Children[index] = left
	parent.Children[index+1] = right

	left.Parent = parent
	right.Parent = parent
}

func fixParents(newParent *Node, children *[]*Node) {
	for _, node := range *children {
		node.Parent = newParent
	}
}
