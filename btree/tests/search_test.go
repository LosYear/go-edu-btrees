package tests

import "testing"
import ".."

func TestBTree_Search_Existing(t *testing.T) {
	tree := fixtureOne()

	// Key 4
	result := tree.Search(4)
	if result == nil || result != tree.Root.Children[0].Children[0] {
		t.Errorf("Key 4 exists in tree but not found")
	}

	// Key 22
	result = tree.Search(22)
	if result == nil || result != tree.Root {
		t.Errorf("Key 22 exists in tree but not found")
	}

	// Key 40
	result = tree.Search(40)
	if result == nil || result != tree.Root.Children[2] {
		t.Errorf("Key 40 exists in tree but not found")
	}

	// Key 31
	result = tree.Search(31)
	if result == nil || result != tree.Root.Children[1].Children[1] {
		t.Errorf("Key 31 exists in tree but not found")
	}
}

func TestBTree_Search_Nonexisting(t *testing.T) {
	tree := fixtureOne()

	// Key 44
	if tree.Search(44) != nil {
		t.Errorf("Key 44 doesn't exists but something found")
	}

	// Key 25
	if tree.Search(27) != nil {
		t.Errorf("Key 27 doesn't exists but something found")
	}

	// Key 0
	if tree.Search(0) != nil {
		t.Errorf("Key 0 doesn't exists but something found")
	}

	// Key 100

	if tree.Search(1000) != nil {
		t.Errorf("Key 0 doesn't exists but something found")
	}
}

func TestBTree_Search_EmptyTree(t *testing.T) {
	tree := btree.BTree{Degree: 2, Root: nil}

	if tree.Search(1) != nil {
		t.Errorf("Tree is empty but something found")
	}
}
