package main

import (
	"fmt"

	rs "github.com/kenshinapa/LetsCodeLeetCode/remove-subdirectories"
)

func main() {
	data := []string{
		"/a/b",
		"/a",
		"/a/b/c",
		"/ax/b",
		"/c/d",
		"/c/d/e",
		"/c",
		"/c/f",
		"/d/d/d",
	}
	result := rs.RemoveSubfolders(data)

	fmt.Println("Result:", result)
}
