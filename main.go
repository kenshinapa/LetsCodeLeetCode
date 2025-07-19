package main

import (
	"fmt"

	add "github.com/kenshinapa/LetsCodeLeetCode/add-two-numbers"
)

func main() {
	result := add.AddTwoNumbers(&add.ListNode{
		Val: 9,
		Next: &add.ListNode{
			Val: 9,
			Next: &add.ListNode{
				Val: 9,
				Next: &add.ListNode{
					Val: 9,
				},
			},
		},
	}, &add.ListNode{
		Val: 9,
		Next: &add.ListNode{
			Val: 9,
			Next: &add.ListNode{
				Val: 9,
				Next: &add.ListNode{
					Val: 9,
				},
			},
		},
	}) // Output: 8 -> 9 -> 9 -> 9 -> 1

	fmt.Println("Result of AddTwoNumbers:")

	// Print first value
	fmt.Printf("%d", result.Val)
	if result.Next != nil {
		fmt.Print(" -> ")
	}

	n := result.Next

	for n != nil {
		fmt.Print(n.Val)
		if n.Next != nil {
			fmt.Print(" -> ")
		}
		n = n.Next
	}
}
