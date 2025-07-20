// Package deleteduplicatefoldersinsystem - Problem 1948. Delete Duplicate Folders in System
package deleteduplicatefoldersinsystem

import (
	"sort"
	"strings"
)

// Due to a bug, there are many duplicate folders in a file system. You are given a 2D array paths, where paths[i] is an array representing an absolute path to the ith folder in the file system.
//
// For example, ["one", "two", "three"] represents the path "/one/two/three".
// Two folders (not necessarily on the same level) are identical if they contain the same non-empty set of identical subfolders and underlying subfolder structure. The folders do not need to be at the root level to be identical. If two or more folders are identical, then mark the folders as well as all their subfolders.
//
// For example, folders "/a" and "/b" in the file structure below are identical. They (as well as their subfolders) should all be marked:
// - /a
// - /a/x
// - /a/x/y
// - /a/z
// - /b
// - /b/x
// - /b/x/y
// - /b/z
// However, if the file structure also included the path "/b/w", then the folders "/a" and "/b" would not be identical. Note that "/a/x" and "/b/x" would still be considered identical even with the added folder.
// Once all the identical folders and their subfolders have been marked, the file system will delete all of them. The file system only runs the deletion once, so any folders that become identical after the initial deletion are not deleted.
//
// Return the 2D array `answer` containing the paths of the remaining folders after deleting all the marked folders. The paths may be returned in any order.

type Tree struct {
	root *folderNode
}

type folderNode struct {
	name                 string
	subfoldersSerialized string
	children             map[string]*folderNode
}

func NewTree() *Tree {
	return &Tree{
		root: &folderNode{
			name:     "",
			children: make(map[string]*folderNode),
		},
	}
}

// Insert adds a new path to the tree.
func (t *Tree) Insert(path []string) {
	current := t.root // Start from the root node

	for _, folder := range path {
		if _, exists := current.children[folder]; !exists {
			current.children[folder] = &folderNode{ // Create a new folder node if it does not exist. Created at `children` map of the current node with the folder name as the key. A single node can have multiple children, each representing a folder in the path.
				name:     folder,
				children: make(map[string]*folderNode),
			}
		}

		current = current.children[folder] // Move to the child node corresponding to the folder
	}
}

func DeleteDuplicateFolder(paths [][]string) [][]string {
	tree := NewTree()

	for _, path := range paths {
		tree.Insert(path)
	}

	freq := make(map[string]int) // hash table records to count the frequency of each serialized representation of the folder structure

	// Post-order traversal based on depth-first search. We calculate the serialized representation of each node structure
	var construct func(node *folderNode)

	// Example input [["a"],["c"],["a","b"],["c","b"],["a","b","x"],["a","b","x","y"],["w"],["w","y"]]

	construct = func(node *folderNode) {
		if len(node.children) == 0 {
			return // If it is a leaf node, return. Leaf nodes do not need to be serialized, and they are located at the end of the path. The actual serialization represents the subfolders of the current node.
		}

		v := make([]string, 0, len(node.children))

		for folder, child := range node.children {
			construct(child)
			v = append(v, folder+"("+child.subfoldersSerialized+")")
		}

		// fmt.Println(v)

		sort.Strings(v)
		// fmt.Println()
		//
		// fmt.Println(v)
		// fmt.Println("-----------------")

		node.subfoldersSerialized = strings.Join(v, "")
		freq[node.subfoldersSerialized]++
	}

	construct(tree.root)

	answer := make([][]string, 0)
	path := make([]string, 0)

	// operate the tree in a depth-first search manner to find the remaining folders after deleting the duplicate folders
	var operate func(node *folderNode)

	operate = func(node *folderNode) {
		// fmt.Println(path)

		if freq[node.subfoldersSerialized] > 1 {
			return // If the frequency of the serialized representation is greater than 1, it means it is a duplicate folder, so we skip it. This means that duplicated folders and their subfoldersSerialized will not be included in the final answer.
		}

		// fmt.Println(path)
		// fmt.Println("------------------")

		if len(path) > 0 {
			tmp := make([]string, len(path))

			copy(tmp, path) // Destination, Source

			answer = append(answer, tmp) // Add the current path to the answer. We need to copy the path slice to avoid modifying it later.
		}

		for folder, child := range node.children { // The children map has the folder name as keys and the child node as values
			path = append(path, folder)

			operate(child)

			path = path[:len(path)-1] // Backtrack by removing the last folder from the path slice
		}
	}

	operate(tree.root)

	// for k, v := range freq {
	// 	if v > 1 {
	// 		fmt.Printf("Duplicate folder structure: %s with frequency %d\n", k, v)
	// 	}
	// }

	return answer
}
