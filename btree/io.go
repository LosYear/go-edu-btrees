package btree

import (
	"fmt"
	"strings"
)

func (tree BTree) Print(withDebugInfo bool) {
	printRoutine(tree.Root, 0, withDebugInfo)
}

func printRoutine(node *Node, level int, withDebugInfo bool) {
	if node == nil {
		fmt.Printf("%s NIL \n", strings.Repeat("–", level+1))
		return
	}

	fmt.Printf("%s %v", strings.Repeat("–", level+1), node.Keys)

	if withDebugInfo {
		fmt.Printf(" | Addr: %p | Parent: %p", node, node.Parent)
	}

	fmt.Printf("\n")

	for _, child := range node.Children {
		printRoutine(child, level+1, withDebugInfo)
	}
}
