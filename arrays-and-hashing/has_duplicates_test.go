package arrays

import (
	"testing"
)

func TestHasDuplicate(t *testing.T) {
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
		// Add more test cases here
	}

	// 3. Iterate over the table and run subtests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validAnagram(tt.args.nums)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// that the core of the problem i need to check to see if a value exists in a slice
// i can iterate over the slice, adding to a map, until i see a map value > 1,
// if that ever becomes true, i return true, else i return false.
// this should be O(n) complexity, where the number of operations scales linearly
// with the the of the array.
func validAnagram(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}
