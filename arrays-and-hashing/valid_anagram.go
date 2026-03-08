// Package arrays has all of the array and hash questions from neetcode
package arrays

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	// 1. Define the input arguments
	type args struct {
		nums []int
	}

	// 2. Define the test table
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			name:     "Example 1: Standard success case",
			args:     args{nums: []int{1, 2, 3, 3}},
			expected: true,
		},
		{
			name:     "Example 2: Standard failure case",
			args:     args{nums: []int{1, 2, 3, 4}},
			expected: false,
		},
	}

	// 3. Iterate over the table and run subtests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasDuplicate(tt.args.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// The core of the problem i need to check to check if a value appears in a slice
// this should be O(n) complexity, where the number of operations scales linearly
// with the size of the slice.
// i am making a map of ints that point to an empty struct, which
// take up no memory, allowing us to use a simple pattern for key checking, that doesnt care about
// the frequency.
func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}
