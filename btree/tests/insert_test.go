package tests

import (
	".."
	"testing"
)


func TestBTree_Insert_NonfullNode(t *testing.T) {
	tree := fixtureOne()

	tree.Insert(5)
	checkTreeConstraits(t, tree)

	node := tree.Root.Children[0].Children[0]

	if len(node.Keys) != 2 || node.Keys[0] != 4 || node.Keys[1] != 5 {
		t.Errorf("Insertion to non-full node failed")
	}
}

func TestBTree_Insert_EmptyTree(t *testing.T) {
	tree := btree.BTree{Degree: 3}
	tree.Insert(5)

	checkTreeConstraits(t, tree)

	if len(tree.Root.Keys) != 1 && tree.Root.Keys[0] != 5 {
		t.Errorf("Insertion to empty tree failed")

	}
}

func TestBTree_Insert_Case1(t *testing.T) {
	tree := btree.BTree{Degree: 3}
	for i := 1; i <= 10; i++ {
		tree.Insert(i)
	}

	checkTreeConstraits(t, tree)

	assertKeysEqual(t, []int{4}, tree.Root.Keys)

	assertKeysEqual(t, []int{2}, tree.Root.Children[0].Keys)
	assertKeysEqual(t, []int{6, 8}, tree.Root.Children[1].Keys)

	assertKeysEqual(t, []int{1}, tree.Root.Children[0].Children[0].Keys)
	assertKeysEqual(t, []int{3}, tree.Root.Children[0].Children[1].Keys)

	assertKeysEqual(t, []int{5}, tree.Root.Children[1].Children[0].Keys)
	assertKeysEqual(t, []int{7}, tree.Root.Children[1].Children[1].Keys)
	assertKeysEqual(t, []int{9, 10}, tree.Root.Children[1].Children[2].Keys)

	// Check parent reverse path
	head := tree.Root
	if !(tree.Root.Children[0].Children[0].Parent.Parent == head &&
		tree.Root.Children[0].Children[1].Parent.Parent == head &&
		tree.Root.Children[1].Children[0].Parent.Parent == head &&
		tree.Root.Children[1].Children[1].Parent.Parent == head &&
		tree.Root.Children[1].Children[2].Parent.Parent == head) {
		t.Errorf("Parent reverse path corrupted")

	}

}

func TestBTree_Insert_Case2(t *testing.T) {
	tree := btree.BTree{Degree: 4}

	numbers := []int{30, 70, 8, 25, 40, 50, 76, 88, 71, 73, 75, 77,
		85, 89, 97, 1, 3, 7, 15, 21, 23, 26, 28, 35, 38, 42, 49, 56, 57}

	for _, n := range numbers {
		tree.Insert(n)
	}

	checkTreeConstraits(t, tree)

	assertKeysEqual(t, []int{30, 76}, tree.Root.Keys)

	assertKeysEqual(t, []int{8, 23}, tree.Root.Children[0].Keys)
	assertKeysEqual(t, []int{40, 50, 70}, tree.Root.Children[1].Keys)
	assertKeysEqual(t, []int{88}, tree.Root.Children[2].Keys)

	assertKeysEqual(t, []int{1, 3, 7}, tree.Root.Children[0].Children[0].Keys)
	assertKeysEqual(t, []int{15, 21}, tree.Root.Children[0].Children[1].Keys)
	assertKeysEqual(t, []int{25, 26, 28}, tree.Root.Children[0].Children[2].Keys)

	assertKeysEqual(t, []int{35, 38}, tree.Root.Children[1].Children[0].Keys)
	assertKeysEqual(t, []int{42, 49}, tree.Root.Children[1].Children[1].Keys)
	assertKeysEqual(t, []int{56, 57}, tree.Root.Children[1].Children[2].Keys)
	assertKeysEqual(t, []int{71, 73, 75}, tree.Root.Children[1].Children[3].Keys)

	assertKeysEqual(t, []int{77, 85}, tree.Root.Children[2].Children[0].Keys)
	assertKeysEqual(t, []int{89, 97}, tree.Root.Children[2].Children[1].Keys)

	// Check parent reverse path
	head := tree.Root
	if !(tree.Root.Children[0].Children[0].Parent.Parent == head &&
		tree.Root.Children[0].Children[1].Parent.Parent == head &&
		tree.Root.Children[0].Children[2].Parent.Parent == head &&

		tree.Root.Children[1].Children[0].Parent.Parent == head &&
		tree.Root.Children[1].Children[1].Parent.Parent == head &&
		tree.Root.Children[1].Children[2].Parent.Parent == head &&
		tree.Root.Children[1].Children[3].Parent.Parent == head &&

		tree.Root.Children[2].Children[0].Parent.Parent == head &&
		tree.Root.Children[2].Children[1].Parent.Parent == head) {
		t.Errorf("Parent reverse path corrupted")

	}
}
