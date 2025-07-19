// Package finduniquebinarystring - Problem 1980. Find Unique Binary String
package finduniquebinarystring

// FindDifferentBinaryString
// Given an array of strings nums containing n unique binary strings each of length n, return a binary string of length n that does not appear in nums. If there are multiple answers, you may return any of them.
// Note: exporting function to import in main.go. In Platform, the function name is unexported.
func FindDifferentBinaryString(nums []string) string {
	n := len(nums)
	res := make([]byte, n)

	for i := 0; i < n; i++ { // Iterate through each index once.
		if nums[i][i] == '0' { // If the character at index i in the string nums[i] is '0', we set res[i] to '1'.
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}

	return string(res)
}
