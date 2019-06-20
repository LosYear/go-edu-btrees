package btree

type Node struct {
	Keys     []int
	Parent   *Node
	Children []*Node
}

type BTree struct {
	Degree uint
	Root   *Node
}