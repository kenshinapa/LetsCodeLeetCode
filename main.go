package main

import (
	"fmt"

	f "github.com/kenshinapa/LetsCodeLeetCode/find-elements-in-contaminated-b-tree"
)

func main() {
	tree := f.Constructor(&f.TreeNode{
		Val: -1,
		Right: &f.TreeNode{
			Val: -1,
		},
	})

	searchOne := tree.Find(1) // Should return false
	searchTwo := tree.Find(2) // Should return true

	fmt.Println("All searches for tree 1 are:", searchOne, searchTwo) // Should return false, true

	tree2 := f.Constructor(&f.TreeNode{
		Val: -1,
		Left: &f.TreeNode{
			Val: -1,
			Left: &f.TreeNode{
				Val: -1,
			},
			Right: &f.TreeNode{
				Val: -1,
			},
		},
		Right: &f.TreeNode{
			Val: -1,
		},
	})

	// 1, 3, 5
	searchOne = tree2.Find(1)    // Should return true
	searchThree := tree2.Find(3) // Should return true
	searchFive := tree2.Find(5)  // Should return false

	fmt.Printf("All seaches for tree 2 are: %v, %v, %v\n", searchOne, searchThree, searchFive) // Should return true, false, true
}
