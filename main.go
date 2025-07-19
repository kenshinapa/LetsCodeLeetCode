package main

import (
	"fmt"

	b "github.com/kenshinapa/LetsCodeLeetCode/find-unique-binary-string"
)

func main() {
	data := []string{
		"001",
		"011",
		"100",
	}
	result := b.FindDifferentBinaryString(data)

	fmt.Println("Result:", result)
}
