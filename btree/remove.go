package btree

import "log"

func (tree *BTree) Remove(key int) {
	node, found := searchRoutine(tree.Root, key)

	if !found {
		log.Fatalln("Key doesn't exists in tree")
	}

	if node.IsLeaf() {
		// If leaf, then simply remove given key and run rebalancing
		node.removeKey(key)

		tree.rebalance(node)

		if len(tree.Root.Keys) == 0 {
			tree.Root = nil
		}
	} else {
		// If internal node, choose new separator (the largest element from left subtree) and swap with deleted key
		// Then rebalance starting from the left "largest" leaf

		keyIndex := node.FindKey(key)

		// Here we won't find actual key, but we'll find the node with the largest child though
		leftLargestNode, _ := searchRoutine(node.Children[keyIndex], key)

		newSeparatorIndex := len(leftLargestNode.Keys) - 1
		newSeparator := leftLargestNode.Keys[newSeparatorIndex]

		node.Keys[keyIndex] = newSeparator

		leftLargestNode.removeKeyByIndex(newSeparatorIndex)

		tree.rebalance(leftLargestNode)
	}
}

func (tree *BTree) rebalance(node *Node) {
	// Check do we need rebalancing
	if node == nil || len(node.Keys) >= tree.minKeysCount() {
		return
	}

	nodeIndex := node.nodeIndex()

	// Borrow elements from right sibling
	rightSibling := node.rightSibling()

	if rightSibling != nil && len(rightSibling.Keys) > tree.minKeysCount() {
		node.Keys = append(node.Keys, node.Parent.Keys[nodeIndex])
		node.Parent.Keys[nodeIndex] = rightSibling.Keys[0]
		rightSibling.removeKeyByIndex(0)

		// We've removed one key, so nodes have to be merged
		if !rightSibling.IsLeaf() {
			rightSibling.Children[0].Parent = node
			node.Children = append(node.Children, rightSibling.Children[0])

			rightSibling.Children[0] = nil
			rightSibling.Children = append([]*Node{}, rightSibling.Children[1:]...)
		}
		return
	}

	// Try borrowing elements from left sibling
	leftSibling := node.leftSibling()

	if leftSibling != nil && len(leftSibling.Keys) > tree.minKeysCount() {
		leftSiblingIndex := nodeIndex - 1

		node.Keys = append([]int{node.Parent.Keys[leftSiblingIndex]}, node.Keys...)
		node.Parent.Keys[leftSiblingIndex] = leftSibling.Keys[len(leftSibling.Keys)-1]
		leftSibling.removeKeyByIndex(len(leftSibling.Keys) - 1)

		// Similarly, we lost one key, merge Children, prepend
		if !leftSibling.IsLeaf() {
			lastIndex := len(leftSibling.Children) - 1
			leftSibling.Children[lastIndex].Parent = node
			node.Children = append([]*Node{leftSibling.Children[lastIndex]}, node.Children...)

			leftSibling.Children[lastIndex] = nil
			leftSibling.Children = append([]*Node{}, leftSibling.Children[:lastIndex]...)
		}

		return
	}

	// merge with siblings if we can't borrow
	if rightSibling != nil {
		// merge with right sibling
		node.Keys = append(node.Keys, node.Parent.Keys[nodeIndex])
		node.Keys = append(node.Keys, rightSibling.Keys...)

		node.Parent.removeKeyByIndex(nodeIndex)
		node.Children = append(node.Children, node.Parent.Children[nodeIndex+1].Children...)
		fixParents(node, &node.Children)

		node.Parent.Children = append(node.Parent.Children[:nodeIndex+1], node.Parent.Children[nodeIndex+2:]...)
	} else if leftSibling != nil {
		// merge with left sibling
		entries := append([]int{}, leftSibling.Keys...)
		entries = append(entries, node.Parent.Keys[nodeIndex-1])
		node.Keys = append(entries, node.Keys...)

		node.Parent.removeKeyByIndex(nodeIndex - 1)
		node.Children = append(node.Parent.Children[nodeIndex-1].Children, node.Children...)
		fixParents(node, &node.Children)

		node.Parent.Children = append(node.Parent.Children[:nodeIndex-1], node.Parent.Children[nodeIndex:]...)
	}

	// if Root is empty, reassign new Root
	if node.Parent == tree.Root && len(tree.Root.Keys) == 0 {
		tree.Root = node
		node.Parent = nil

		return
	}

	tree.rebalance(node.Parent)
}
