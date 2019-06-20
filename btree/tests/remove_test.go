package tests

import (
	".."
	"testing"
)

func TestBTree_Remove_LeafAndRoot(t *testing.T) {
	t.Run("Remove one element from Root, Root is tree", func(t *testing.T) {
		tree := btree.BTree{Degree: 3}
		tree.Insert(1)
		tree.Insert(2)

		tree.Remove(2)

		checkTreeConstraits(t, tree)

		if len(tree.Root.Keys) != 1 || tree.Root.Keys[0] != 1 {
			t.Errorf("Tree structure is not correct after deletion")
		}
	})

	t.Run("Remove all elements from Root, make tree empty", func(t *testing.T) {
		tree := btree.BTree{Degree: 3}
		tree.Insert(1)
		tree.Insert(13)

		tree.Remove(1)
		tree.Remove(13)

		checkTreeConstraits(t, tree)

		if tree.Root != nil {
			t.Errorf("Tree structure is not correct after deletion, tree have to be empty")
		}
	})
}

func TestBTree_Remove_InternalNode(t *testing.T) {
	t.Run("Remove from internal node, no rebalancing", func(t *testing.T) {
		tree := fixtureTwo()
		tree.Remove(76)

		if len(tree.Root.Children[2].Keys) != 2 || tree.Root.Children[2].Keys[0] != 75 || tree.Root.Children[2].Keys[1] != 88 ||
			len(tree.Root.Children[2].Children[0].Keys) != 2 || tree.Root.Children[2].Children[0].Keys[0] != 71 || tree.Root.Children[2].Children[0].Keys[1] != 73 {
			t.Errorf("Tree structure is not correct after deletion from internal node, no rebalancing needed")
		}
	})

	t.Run("Left rotation, leaf", func(t *testing.T) {
		tree := fixtureTwo()
		tree.Remove(35)
		tree.Remove(38)

		checkTreeConstraits(t, tree)

		checkKeys(t, tree.Root.Children[1], 42, 50)
		checkKeys(t, tree.Root.Children[1].Children[0], 40)
		checkKeys(t, tree.Root.Children[1].Children[1], 49)
		checkKeys(t, tree.Root.Children[1].Children[2], 56, 67)
	})

	t.Run("Right rotation, leaf", func(t *testing.T) {
		tree := fixtureTwo()
		tree.Remove(56)
		tree.Remove(67)

		checkTreeConstraits(t, tree)

		checkKeys(t, tree.Root.Children[1], 40, 49)
		checkKeys(t, tree.Root.Children[1].Children[0], 35, 38)
		checkKeys(t, tree.Root.Children[1].Children[1], 42)
		checkKeys(t, tree.Root.Children[1].Children[2], 50)
	})

	t.Run("No rotation, merging with right sibling", func(t *testing.T) {
		tree := fixtureTwo()
		tree.Remove(42)
		tree.Remove(38)
		tree.Remove(35)

		checkTreeConstraits(t, tree)

		checkKeys(t, tree.Root.Children[1], 50)
		checkKeys(t, tree.Root.Children[1].Children[0], 40, 49)
		checkKeys(t, tree.Root.Children[1].Children[1], 56, 67)
	})

	t.Run("No rotation, merging with left sibling", func(t *testing.T) {
		tree := fixtureTwo()
		tree.Remove(49)
		tree.Remove(56)
		tree.Remove(67)

		checkTreeConstraits(t, tree)

		checkKeys(t, tree.Root.Children[1], 40)
		checkKeys(t, tree.Root.Children[1].Children[0], 35, 38)
		checkKeys(t, tree.Root.Children[1].Children[1], 42, 50)

	})

	t.Run("Remove from internal node, complex case", func(t *testing.T) {
		tree := fixtureThree()
		tree.Remove(25)
		tree.Remove(23)
		checkTreeConstraits(t, tree)

		checkKeys(t, tree.Root.Children[0], 21, 40)
		checkKeys(t, tree.Root.Children[0].Children[1], 30)
		checkKeys(t, tree.Root.Children[0].Children[1].Children[0], 26, 28)
		checkKeys(t, tree.Root.Children[0].Children[1].Children[1], 35, 38)
		checkKeys(t, tree.Root.Children[0].Children[2], 50)
		checkKeys(t, tree.Root.Children[0].Children[2].Children[0], 42, 49)
		checkKeys(t, tree.Root.Children[0].Children[2].Children[1], 56, 67)

		tree.Remove(100)
		tree.Remove(101)
		tree.Remove(102)
		tree.Remove(103)
		tree.Remove(110)
		tree.Remove(105)
		checkTreeConstraits(t, tree)

		checkKeys(t, tree.Root.Children[1], 88)
		checkKeys(t, tree.Root.Children[1].Children[0], 76)
		checkKeys(t, tree.Root.Children[1].Children[1], 99)
		checkKeys(t, tree.Root.Children[1].Children[1], 99)
		checkKeys(t, tree.Root.Children[1].Children[0].Children[0], 71, 73, 75)
		checkKeys(t, tree.Root.Children[1].Children[0].Children[1], 77, 85)
		checkKeys(t, tree.Root.Children[1].Children[1].Children[0], 89, 97)
		checkKeys(t, tree.Root.Children[1].Children[1].Children[1], 135, 140)
	})
}

func TestBTree_Remove_HeightReducing(t *testing.T) {
	tree := fixtureTwo()
	checkTreeConstraits(t, tree)
	tree.Remove(1)
	checkTreeConstraits(t, tree)
	tree.Remove(3)
	checkTreeConstraits(t, tree)
	tree.Remove(7)
	checkTreeConstraits(t, tree)
	tree.Remove(8)
	checkTreeConstraits(t, tree)
	tree.Remove(15)
	checkTreeConstraits(t, tree)
	tree.Remove(23)
	checkTreeConstraits(t, tree)
	tree.Remove(25)
	checkTreeConstraits(t, tree)
	tree.Remove(26)
	checkTreeConstraits(t, tree)
	tree.Remove(21)
	checkTreeConstraits(t, tree)
	tree.Remove(28)
	checkTreeConstraits(t, tree)
	tree.Remove(30)
	checkTreeConstraits(t, tree)
	tree.Remove(89)
	checkTreeConstraits(t, tree)
	tree.Remove(97)
	checkTreeConstraits(t, tree)
	tree.Remove(88)
	checkTreeConstraits(t, tree)
	tree.Remove(77)
	checkTreeConstraits(t, tree)
	tree.Remove(76)
	checkTreeConstraits(t, tree)
	tree.Remove(85)
	checkTreeConstraits(t, tree)
	tree.Remove(71)
	checkTreeConstraits(t, tree)
	tree.Remove(50)
	checkTreeConstraits(t, tree)
	tree.Remove(42)
	checkTreeConstraits(t, tree)
	tree.Remove(67)
	checkTreeConstraits(t, tree)
	checkKeys(t, tree.Root, 49)
	checkKeys(t, tree.Root.Children[0], 38)
	checkKeys(t, tree.Root.Children[0].Children[0], 35)
	checkKeys(t, tree.Root.Children[0].Children[1], 40)
	checkKeys(t, tree.Root.Children[1], 70)
	checkKeys(t, tree.Root.Children[1].Children[0], 56)
	checkKeys(t, tree.Root.Children[1].Children[1], 73, 75)

	tree.Remove(40)
	checkTreeConstraits(t, tree)
	checkKeys(t, tree.Root, 49, 70)
	checkKeys(t, tree.Root.Children[0], 35, 38)
	checkKeys(t, tree.Root.Children[1], 56)
	checkKeys(t, tree.Root.Children[2], 73, 75)
}
