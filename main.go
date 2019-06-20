package main

import (
	"./btree"
	"fmt"
)

func degree3Insertion() btree.BTree {
	tree := btree.BTree{Degree: 3}

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(8)
	tree.Insert(9)
	tree.Insert(10)

	tree.Remove(4)

	return tree
}

func degree4Insertiob() btree.BTree {
	tree := btree.BTree{Degree: 4}

	numbers := []int{30, 70, 8, 25, 40, 50, 76, 88, 71, 73, 75, 77,
		85, 89, 97, 1, 3, 7, 15, 21, 23, 26, 28, 35, 38, 42, 49, 56, 57}

	til := 29

	for _, n := range numbers[:til] {
		tree.Insert(n)

	}

	return tree
}

func removeLeafAndRoot() btree.BTree {
	tree := btree.BTree{Degree: 3}
	tree.Insert(1)
	tree.Insert(2)

	tree.Remove(2)

	return tree
}

func complexRemove() {
	tree := degree4Insertiob()
	fmt.Println("INITIAL:")
	tree.Print(false)
	fmt.Println(" ")

	keys := []int{1, 3, 7, 8, 15, 23, 25, 26, 21, 28, 30, 89, 97, 88, 77, 76, 85, 71, 50, 42, 56, 40}

	for _, key := range keys {
		fmt.Println("REMOVED ", key)
		tree.Remove(key)
		tree.Print(false)
		fmt.Println(" ")
	}

}

func main() {
	degree3Insertion().Print(false)
	//degree4Insertiob().Print(false)

	//removeLeafAndRoot().Print(false)
	//complexRemove()
}
