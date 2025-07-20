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

// Notes/insights:
// To be honest, I had to look up the solution. I was not aware about serialization, but I was able to build the
// structure correctly. Only that just with a name instead of the subfolderSerialization thing. To be honest, I think
// that if someone is not aware of the existence of serialization considering only the children, this would be much
// harder to implement. I did think about building strings like /a/b/x for "a", "b", "x" and then sort it and try to
// study the strings searching for repetitions. But it seems more efficient with the serialization. I see that
// solving this is not just about knowing how to implement a tree structure and insert stuff into it. It is about
// thinking outside the box. It is about realizing that the metadata included on each node inside the tree is flexible.
// We can include stuff that could help us solve a problem even if it is not part of a regular tree structure where we
// only care about linking nodes. In a previous problem, we used a map for visited nodes at the root of the tree.
// This way, searches were not needed once the tree is completed. And we can return answers to see if a given value
// exists on the tree in linear time. I learned that we can add fields to a node in order to work for us in our benefit
// when solving a bigger problem.

// Copilot response (yes, I like to discuss insights with it):
// That’s a great insight! Serialization is a powerful technique for comparing complex structures like trees, especially when you need to detect duplicates based on structure rather than just values or paths. By storing extra metadata (like subfoldersSerialized) in each node, you can efficiently solve problems that go beyond basic tree operations.
//
// Your reflection is correct:
//
//
// Tree nodes can hold any data that helps solve the problem, not just names or pointers.
// Serialization lets you compare subtrees quickly and accurately.
// Using maps for fast lookup or frequency counting is a common optimization.
// This approach—augmenting data structures with problem-specific fields—often leads to more efficient and elegant solutions. It’s a key skill in algorithm design.

// Me again:
// One more insight. I think that if we included the current leaf in the serialization it would not be super useful
// since on the operate function we would still have to check on its children. We want to have a peek into the future
// through the serialization, right? If so, I think that it is not only adding the metadata that can help us but also
// about building it correctly.

// Copilot response:
// Correct. The power of serialization here is that it encodes the structure of a node’s children, allowing you to
// identify duplicate subtrees without needing to check each child individually during the operate phase. Including the
// current node’s name in its own serialization would not help, because duplication is determined by the structure
// beneath each node, not the node itself.
//
// So, it’s not just about adding metadata—it’s about designing that metadata to capture exactly what you need for
// efficient duplicate detection. Serialization gives you a “preview” of the subtree, enabling quick frequency checks
// and pruning. Building the metadata correctly is key to solving the problem efficiently.
