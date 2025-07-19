// Package regularbinarytree is not part of LeetCode problems. I just want to remember how to implement a regular binary tree in Go.
package regularbinarytree

import "strconv"

type TreeNode struct {
	Val   int64
	Left  *TreeNode
	Right *TreeNode
}

type RegularBinaryTree struct {
	root *TreeNode
}

func (rbt *RegularBinaryTree) JsonString() string {
	if rbt.root == nil {
		return "Empty Tree"
	}

	return rbt.jsonString(rbt.root)
}

func (rbt *RegularBinaryTree) jsonString(node *TreeNode) string {
	if node == nil {
		return "null"
	}

	return "{\"Val\": " + strconv.FormatInt(node.Val, 10) + ", \"Left\": " + rbt.jsonString(node.Left) + ", \"Right\": " + rbt.jsonString(node.Right) + "}"
}

func NewRegularBinaryTree(root *TreeNode) *RegularBinaryTree {
	return &RegularBinaryTree{
		root: root,
	}
}

func (rbt *RegularBinaryTree) Insert(value int64) {
	if rbt.root == nil {
		rbt.root = &TreeNode{Val: value}
		return
	}

	rbt.insertRecursive(rbt.root, value)
}

func (rbt *RegularBinaryTree) insertRecursive(node *TreeNode, value int64) {
	if value < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{Val: value}
		} else {
			rbt.insertRecursive(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Val: value}
		} else {
			rbt.insertRecursive(node.Right, value)
		}
	}
}

func (rbt *RegularBinaryTree) Search(value int64) bool {
	return rbt.searchRecursive(rbt.root, value)
}

func (rbt *RegularBinaryTree) searchRecursive(node *TreeNode, value int64) bool {
	if node == nil {
		return false
	}

	if node.Val == value {
		return true
	}

	if value < node.Val {
		return rbt.searchRecursive(node.Left, value)
	}

	return rbt.searchRecursive(node.Right, value)
}

func (rbt *RegularBinaryTree) InOrderTraversal() []int64 {
	result := []int64{}

	rbt.inOrderRecursive(rbt.root, &result)

	return result
}

func (rbt *RegularBinaryTree) inOrderRecursive(node *TreeNode, result *[]int64) {
	if node == nil {
		return
	}

	rbt.inOrderRecursive(node.Left, result)

	*result = append(*result, node.Val)

	rbt.inOrderRecursive(node.Right, result)
}

func (rbt *RegularBinaryTree) PreOrderTraversal() []int64 {
	result := []int64{}

	rbt.preOrderRecursive(rbt.root, &result)

	return result
}

func (rbt *RegularBinaryTree) preOrderRecursive(node *TreeNode, result *[]int64) {
	if node == nil {
		return
	}

	*result = append(*result, node.Val)

	rbt.preOrderRecursive(node.Left, result)
	rbt.preOrderRecursive(node.Right, result)
}

func (rbt *RegularBinaryTree) Delete(value int64) {
	rbt.root = rbt.deleteRecursive(rbt.root, value)
}

func (rbt *RegularBinaryTree) deleteRecursive(node *TreeNode, value int64) *TreeNode {
	if node == nil {
		return nil
	}

	if value < node.Val {
		node.Left = rbt.deleteRecursive(node.Left, value)
	} else if value > node.Val {
		node.Right = rbt.deleteRecursive(node.Right, value)
	} else {
		// Node with one child or no child
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Node with two children: Get the inorder successor (smallest in the right subtree)
		minNode := rbt.minValueNode(node.Right)
		node.Val = minNode.Val
		node.Right = rbt.deleteRecursive(node.Right, minNode.Val)
	}

	return node
}

func (rbt *RegularBinaryTree) minValueNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func (rbt *RegularBinaryTree) MaxValue() int64 {
	if rbt.root == nil {
		return 0 // or some other value indicating the tree is empty
	}

	node := rbt.root
	for node.Right != nil {
		node = node.Right
	}

	return node.Val
}

func (rbt *RegularBinaryTree) MinValue() int64 {
	if rbt.root == nil {
		return 0 // or some other value indicating the tree is empty
	}

	node := rbt.root
	for node.Left != nil {
		node = node.Left
	}

	return node.Val
}

func (rbt *RegularBinaryTree) Height() int {
	return rbt.heightRecursive(rbt.root)
}

func (rbt *RegularBinaryTree) heightRecursive(node *TreeNode) int {
	if node == nil {
		return -1 // Return -1 for height of empty tree
	}

	leftHeight := rbt.heightRecursive(node.Left)
	rightHeight := rbt.heightRecursive(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func (rbt *RegularBinaryTree) IsBalanced() bool {
	return rbt.isBalancedRecursive(rbt.root) != -1
}

func (rbt *RegularBinaryTree) isBalancedRecursive(node *TreeNode) int {
	if node == nil {
		return 0 // Height of empty tree is 0
	}

	leftHeight := rbt.isBalancedRecursive(node.Left)
	if leftHeight == -1 {
		return -1 // Left subtree is not balanced
	}

	rightHeight := rbt.isBalancedRecursive(node.Right)
	if rightHeight == -1 {
		return -1 // Right subtree is not balanced
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1 // Current node is not balanced
	}

	return max(leftHeight, rightHeight) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (rbt *RegularBinaryTree) PrettyPrint() string {
	if rbt.root == nil {
		return "Empty Tree"
	}

	return prettyPrintNode(rbt.root, "", true)
}

func prettyPrintNode(node *TreeNode, prefix string, isTail bool) string {
	if node == nil {
		return ""
	}

	result := ""

	if node.Right != nil {
		result += prettyPrintNode(node.Right, prefix+ifThenElse(isTail, "    ", "│   "), false)
	}

	result += prefix

	if isTail {
		result += "└── "
	} else {
		result += "┌── "
	}

	result += strconv.FormatInt(node.Val, 10) + "\n"

	if node.Left != nil {
		result += prettyPrintNode(node.Left, prefix+ifThenElse(isTail, "    ", "│   "), true)
	}

	return result
}

func ifThenElse(cond bool, a, b string) string {
	if cond {
		return a
	}

	return b
}

func (rbt *RegularBinaryTree) Balance() {
	if rbt.root == nil {
		return
	}
	rbt.root = rbt.balanceRecursive(rbt.root)
}

func (rbt *RegularBinaryTree) balanceRecursive(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	node.Left = rbt.balanceRecursive(node.Left)
	node.Right = rbt.balanceRecursive(node.Right)

	leftHeight := rbt.heightRecursive(node.Left)
	rightHeight := rbt.heightRecursive(node.Right)

	if abs(leftHeight-rightHeight) <= 1 {
		return node // Already balanced
	}

	if leftHeight > rightHeight { // Left subtree is taller, rotate right
		return rbt.rotateRight(node)
	} else { // Right subtree is taller, rotate left
		return rbt.rotateLeft(node)
	}
}

func (rbt *RegularBinaryTree) rotateRight(node *TreeNode) *TreeNode {
	if node == nil || node.Left == nil {
		return node // Cannot rotate right if node or left child is nil
	}

	newRoot := node.Left

	node.Left = newRoot.Right
	newRoot.Right = node

	return newRoot
}

func (rbt *RegularBinaryTree) rotateLeft(node *TreeNode) *TreeNode {
	if node == nil || node.Right == nil {
		return node // Cannot rotate left if node or right child is nil
	}

	newRoot := node.Right

	node.Right = newRoot.Left
	newRoot.Left = node

	return newRoot
}

func (rbt *RegularBinaryTree) InsertAndBalance(value int64) {
	rbt.Insert(value)
	rbt.Balance()
}

func (rbt *RegularBinaryTree) DeleteAndBalance(value int64) {
	rbt.Delete(value)
	rbt.Balance()
}
