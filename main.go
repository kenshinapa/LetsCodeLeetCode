package main

import (
	"fmt"

	t "github.com/kenshinapa/LetsCodeLeetCode/regular-binary-tree"
)

func main() {
	tree := t.NewRegularBinaryTree(&t.TreeNode{Val: 10})

	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(18)
	tree.Insert(1)
	tree.Insert(0)
	tree.Insert(8)
	tree.InsertAndBalance(20)

	tree.Delete(5)

	fmt.Println(tree.PrettyPrint())
	fmt.Println(tree.IsBalanced())
}
