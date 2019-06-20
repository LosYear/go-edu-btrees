package btree

func (node *Node) FindKey(value int) int {
	left, right := -1, len(node.Keys)

	for left < right-1 {
		middle := (left + right) / 2
		if node.Keys[middle] < value {
			left = middle
		} else {
			right = middle
		}

	}

	return right
}

func (node *Node) removeKey(key int) {
	index := node.FindKey(key)

	node.removeKeyByIndex(index)
}

func (node *Node) removeKeyByIndex(index int) {
	node.Keys = append(node.Keys[:index], node.Keys[index+1:]...)
}

// Keep in mind, that function just inserts key into slice without checking for tree constraints
func (node *Node) insertKey(key int) {
	index := node.FindKey(key)

	node.Keys = append(node.Keys, 0)
	copy(node.Keys[index+1:], node.Keys[index:])

	node.Keys[index] = key
}

func (node *Node) nodeIndex() int {
	// We assume that Parent exists
	parent := node.Parent

	if parent != nil {
		for index, child := range parent.Children {
			if child == node {
				return index
			}
		}
	}

	return -1
}

// Returns left sibling of given node if any
func (node *Node) leftSibling() *Node {
	if node.Parent == nil {
		return nil
	}

	nodeIndex := node.nodeIndex()

	if nodeIndex != 0 {
		return node.Parent.Children[nodeIndex-1]
	}

	return nil
}

func (node *Node) rightSibling() *Node {
	if node.Parent == nil {
		return nil
	}

	nodeIndex := node.nodeIndex()

	if nodeIndex < len(node.Parent.Children)-1 {
		return node.Parent.Children[nodeIndex+1]
	}
	return nil
}
