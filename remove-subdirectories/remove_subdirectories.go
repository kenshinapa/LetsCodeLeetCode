// Package removesubdirectories - Problem 1233. Remove Sub-Folders from the Filesystem
package removesubdirectories

// RemoveSubfolders
// Given a list of folders folder, return the folders after removing all sub-folders in those folders. You may return the answer in any order.
//
// If a folder[i] is located within another folder[j], it is called a sub-folder of it. A sub-folder of folder[j] must start with folder[j], followed by a "/". For example, "/a/b" is a sub-folder of "/a", but "/b" is not a sub-folder of "/a/b/c".
//
// The format of a path is one or more concatenated strings of the form: '/' followed by one or more lowercase English letters.
// Note: exporting function to import in main.go. In Platform, the function name is unexported.
func RemoveSubfolders(folder []string) []string { //
	folderMap := make(map[string]struct{})
	for _, f := range folder {
		folderMap[f] = struct{}{} // Initialize the map with all folders
	}

	var result []string

	for _, f := range folder {
		_, ok := folderMap[f]
		if !ok {
			continue
		}

		for fm := range folderMap { // Possible subfolder
			if f == fm {
				continue
			}

			if len(fm) > len(f) && fm[:len(f)] == f && fm[len(f)] == '/' {
				delete(folderMap, fm)
			}
		}
	}

	for f := range folderMap { // Iterate through the map to collect valid folders
		result = append(result, f)
	}

	return result
}
