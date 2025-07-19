// Package removesubdirectories - Problem 1233. Remove Sub-Folders from the Filesystem
package removesubdirectories

import (
	"sort"
	"strings"
)

// RemoveSubfolders
// Given a list of folders folder, return the folders after removing all sub-folders in those folders. You may return the answer in any order.
//
// If a folder[i] is located within another folder[j], it is called a sub-folder of it. A sub-folder of folder[j] must start with folder[j], followed by a "/". For example, "/a/b" is a sub-folder of "/a", but "/b" is not a sub-folder of "/a/b/c".
//
// The format of a path is one or more concatenated strings of the form: '/' followed by one or more lowercase English letters.
// Note: exporting function to import in main.go. In Platform, the function name is unexported.
func RemoveSubfolders(folder []string) []string { //
	if len(folder) == 0 {
		return nil
	}

	sorted := make([]string, len(folder))
	copy(sorted, folder)
	// Sort lexicographically
	sort.Strings(sorted)

	var result []string
	for i, f := range sorted {
		if i == 0 || !strings.HasPrefix(f, result[len(result)-1]+"/") {
			result = append(result, f)
		}
	}
	return result
}
