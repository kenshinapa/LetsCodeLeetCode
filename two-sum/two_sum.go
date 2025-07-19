// Package twosum - Problem 1. Two Sum
package twosum

// TwoSum - Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.
func TwoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	return twoSumRecursive(nums, target, 0, make(map[int]int))
}

// example:
// idx:  0,  1,  2, 3
// []int{7, 11, 15, 2}
// target: 9
// output: [0, 3]

func twoSumRecursive(nums []int, target int, index int, visited map[int]int) []int {
	complement := target - nums[index]            // 2 = 9 - 7 | index 1: -2 = 9 - 11 | index 2: -6 = 9 - 15 | index 3: 7 = 9 - 2
	if idx, found := visited[complement]; found { // first, visited is empty
		return []int{idx, index}
	}

	visited[nums[index]] = index // visited: {7: 0} | {7: 0, 11: 1} | {7: 0, 11: 1, 15: 2} | {7: 0, 11: 1, 15: 2, 2: 3}

	return twoSumRecursive(nums, target, index+1, visited)
}
