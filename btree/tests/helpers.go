package tests

import (
	".."
	"reflect"
	"testing"
)

func assertKeysEqual(t *testing.T, expected, got []int) {
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v got %v", expected, got)
	}
}

func checkKeys(t *testing.T, node *btree.Node, keys ...int) {
	if len(node.Keys) != len(keys) {
		t.Errorf("Keys count doesn't match for given node: %v", *node)
		return
	}

	for index, key := range keys {
		if node.Keys[index] != key {
			t.Errorf("Wrong key at index %d, got: %v, expected: %v", index, node.Keys, keys)
			return
		}
	}

	if len(node.Children) > 0 && len(node.Children) != len(node.Keys)+1 {
		t.Errorf("Wrong Children count, k: %v,  c: %v", node.Keys, node.Children)
		return
	}

	for _, child := range node.Children {
		if child.Parent != node {
			t.Errorf("Wrong Parent for child %v", child)
			return
		}
	}
}

func checkParents(t *testing.T, tree *btree.BTree) {
	if tree.Root == nil {
		return
	}

	queue := make([]*btree.Node, 0)

	queue = append(queue, tree.Root)

	for ; len(queue) > 0; {
		head := queue[0]
		queue = queue[1:]

		for _, child := range head.Children {
			if child.Parent != head {
				t.Errorf("Wring Parent for child %v", child)
				return
			}

			queue = append(queue, child)
		}
	}
}

func checkTreeConstraits(t *testing.T, tree btree.BTree) {
	if tree.Root == nil {
		return
	}

	// Check all leaves on the same level
	leaves := make([]*btree.Node, 0)
	nonRootNodes := make([]*btree.Node, 0)
	allNodes := make([]*btree.Node, 0)
	queue := make([]*btree.Node, 0)
	queue = append(queue, tree.Root)

	for ; len(queue) > 0; {
		head := queue[0]
		queue = queue[1:]

		allNodes = append(allNodes, head)

		if head.IsLeaf() {
			leaves = append(leaves, head)
		} else {
			for _, child := range head.Children {
				queue = append(queue, child)
			}
		}
	}

	depth := -1

	for _, leaf := range leaves {
		curDepth := 0
		for current := leaf; current.Parent != nil; current = current.Parent {
			curDepth++
		}

		if depth == -1 {
			depth = curDepth
		} else if curDepth != depth {
			t.Errorf("Leaves must be at the same level")
			return
		}
	}

	if !tree.Root.IsLeaf() && len(tree.Root.Children) < 2 {
		t.Errorf("If the root node is a non leaf node, then it must have at least 2 children.")
		return
	}

	for _, node := range nonRootNodes {
		//  All nodes except root must have at least [m/2]-1 keys and maximum of m-1 keys.
		if len(node.Keys) > int(tree.Degree-1) || len(node.Keys) < int(tree.Degree/2)-1 {
			t.Errorf("All nodes except root must have at least [m/2]-1 keys and maximum of m-1 keys")
			return
		}

		// All non leaf nodes except root (i.e. all internal nodes) must have at least m/2 children.
		if !node.IsLeaf() && len(node.Children) < int(tree.Degree/2) {
			t.Errorf("All non leaf nodes except root (i.e. all internal nodes) must have at least m/2 children.")
			return

		}
	}

	for _, node := range allNodes {
		if !node.IsLeaf() && len(node.Children) != len(node.Keys)+1 {
			t.Errorf("A non leaf node with n-1 keys must have n number of children.")
			return
		}

		for _, child := range node.Children {
			if child.Parent != node {
				t.Errorf("Wrong parent link")
				return
			}
		}
	}
}
