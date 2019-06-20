package tests

import ".."

// Represents this tree
// https://cdn-images-1.medium.com/max/1600/1*lSx54zw4tzOOE09uJdgltg.jpeg
func fixtureOne() btree.BTree {
	node4 := btree.Node{Keys: []int{4}, Children: []*btree.Node{}}
	node815 := btree.Node{Keys: []int{8, 15}, Children: []*btree.Node{}}
	node6 := btree.Node{Keys: []int{6}, Children: []*btree.Node{&node4, &node815}}
	node4.Parent = &node6
	node815.Parent = &node6

	node25 := btree.Node{Keys: []int{25}, Children: []*btree.Node{}}
	node31 := btree.Node{Keys: []int{31}, Children: []*btree.Node{}}
	node29 := btree.Node{Keys: []int{29}, Children: []*btree.Node{&node25, &node31}}
	node31.Parent = &node29
	node25.Parent = &node29

	node39 := btree.Node{Keys: []int{39}, Children: []*btree.Node{}}
	node4850 := btree.Node{Keys: []int{48, 50}, Children: []*btree.Node{}}
	node40 := btree.Node{Keys: []int{40}, Children: []*btree.Node{&node39, &node4850}}
	node39.Parent = &node40
	node4850.Parent = &node40

	node2236 := btree.Node{Keys: []int{22, 36}, Children: []*btree.Node{&node6, &node29, &node40}}
	node40.Parent = &node2236
	node29.Parent = &node2236
	node6.Parent = &node2236

	return btree.BTree{3, &node2236}
}

func fixtureTwo() btree.BTree {
	node137 := btree.Node{Keys: []int{1, 3, 7}}
	node152123 := btree.Node{Keys: []int{15, 21, 23}}
	node2628 := btree.Node{Keys: []int{26, 28}}

	node825 := btree.Node{Keys: []int{8, 25}, Children: []*btree.Node{&node137, &node152123, &node2628}}
	node137.Parent = &node825
	node152123.Parent = &node825
	node2628.Parent = &node825

	node3538 := btree.Node{Keys: []int{35, 38}}
	node4249 := btree.Node{Keys: []int{42, 49}}
	node5667 := btree.Node{Keys: []int{56, 67}}

	node4050 := btree.Node{Keys: []int{40, 50}, Children: []*btree.Node{&node3538, &node4249, &node5667}}
	node3538.Parent = &node4050
	node4249.Parent = &node4050
	node5667.Parent = &node4050

	node717375 := btree.Node{Keys: []int{71, 73, 75}}
	node7785 := btree.Node{Keys: []int{77, 85}}
	node8997 := btree.Node{Keys: []int{89, 97}}

	node7688 := btree.Node{Keys: []int{76, 88}, Children: []*btree.Node{&node717375, &node7785, &node8997}}
	node717375.Parent = &node7688
	node7785.Parent = &node7688
	node8997.Parent = &node7688

	node3070 := btree.Node{Keys: []int{30, 70}, Children: []*btree.Node{&node825, &node4050, &node7688}}
	node825.Parent = &node3070
	node4050.Parent = &node3070
	node7688.Parent = &node3070

	return btree.BTree{Degree: 4, Root: &node3070}
}

func fixtureThree() btree.BTree{
	tree := fixtureTwo()
	tree.Insert(105)
	tree.Insert(110)
	tree.Insert(99)
	tree.Insert(100)
	tree.Insert(101)
	tree.Insert(102)
	tree.Insert(103)
	tree.Insert(135)
	tree.Insert(140)
	tree.Insert(9)
	tree.Insert(10)
	tree.Insert(11)

	return tree
}