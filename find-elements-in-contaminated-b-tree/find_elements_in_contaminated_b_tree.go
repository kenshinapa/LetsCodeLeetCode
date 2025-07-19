// Package findelementsincontaminatedbtree - Problem 1261. Find Elements in a Contaminated Binary Tree
package findelementsincontaminatedbtree

// Given a binary tree with the following rules:
//
// root.val == 0
// For any treeNode:
// - If treeNode.val has a value x and treeNode.left != null, then treeNode.left.val == 2 * x + 1
// - If treeNode.val has a value x and treeNode.right != null, then treeNode.right.val == 2 * x + 2
// Now the binary tree is contaminated, which means all treeNode.val have been changed to -1.
//
// Implement the FindElements class: (Classes in Go, lol)
//
// FindElements(TreeNode* root) Initializes the object with a contaminated binary tree and recovers it.
// - bool find(int target) Returns true if the target value exists in the recovered binary tree.
// - Note: exporting function to import in main.go. In Platform, the function name is unexported.

// FindElements - Definition for a binary tree node.
//
//	type TreeNode struct {
//		Val int
//		Left *TreeNode
//		Right *TreeNode
//	}
type FindElements struct {
	root    *TreeNode
	visited map[int]bool // To keep track of visited nodes
}

// TreeNode - Definition for a binary tree node. I guess this is already implemented in the platform.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	fe := FindElements{
		root:    root,
		visited: make(map[int]bool),
	}

	fe.recoverTree(root, 0)

	return fe
}

func (this *FindElements) recoverTree(node *TreeNode, value int) {
	if node == nil {
		return
	}

	node.Val = value
	this.visited[value] = true

	this.recoverTree(node.Left, 2*value+1)
	this.recoverTree(node.Right, 2*value+2)
}

func (this *FindElements) Find(target int) bool {
	_, exists := this.visited[target]
	return exists
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
