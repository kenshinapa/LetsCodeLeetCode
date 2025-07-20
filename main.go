package main

import (
	"fmt"

	f "github.com/kenshinapa/LetsCodeLeetCode/delete-duplicate-folders-in-system"
)

func main() {
	// Example input [["a"],["c"],["a","b"],["c","b"],["a","b","x"],["a","b","x","y"],["w"],["w","y"]]
	result := f.DeleteDuplicateFolder([][]string{
		{"a"},
		{"c"},
		{"a", "b"},
		{"c", "b"},
		{"a", "b", "x"},
		{"a", "b", "x", "y"},
		{"r", "w"},
		{"r", "w", "y"},
	})

	fmt.Println(result)
}
