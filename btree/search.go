package btree

// Searches for node with given key in BTree
// If key found - returns node itself
// Else - nil
func (tree BTree) Search(key int) *Node {
	node, found :=  searchRoutine(tree.Root, key)

	if found {
		return node
	}

	return nil
}

// Returns node and true if found
// Returns leaf and false if not found (used for insertion)
func searchRoutine(root *Node, key int) (*Node, bool) {
	if root == nil {
		return nil, false
	}

	// Check if key in current node
	keyIndex := root.FindKey(key)

	// Key in current Root
	if keyIndex < len(root.Keys) && root.Keys[keyIndex] == key {
		return root, true
	}

	// Key can be in one of Children nodes
	if keyIndex < len(root.Children) {
		return searchRoutine(root.Children[keyIndex], key)
	}

	return root, false
}