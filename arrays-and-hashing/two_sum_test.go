// twoSum iterates through the array once, storing numbers and their indices in a map.
// For each number, it checks if the complement (target - num) already exists in the map.
// If found, it returns the indices.
// Time Complexity: O(n) - Single pass through the array.
// Space Complexity: O(n) - Map stores at most n elements.
package arrays

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name:     "Example 1: Adjacent elements",
			args:     args{nums: []int{3, 4, 5, 6}, target: 7},
			expected: []int{0, 1},
		},
		{
			name:     "Example 2: Separated elements",
			args:     args{nums: []int{4, 5, 6}, target: 10},
			expected: []int{0, 2},
		},
		{
			name:     "Example 3: No match found",
			args:     args{nums: []int{1, 2, 3}, target: 10},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.args.nums, tt.args.target)
			if !slices.Equal(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func twoSum(nums []int, target int) []int {
	// Map to store the value and its original index: map[value]index
	numToIndex := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		// If the complement exists, return its stored index and the current index
		if matchIndex, ok := numToIndex[complement]; ok {
			return []int{matchIndex, i}
		}

		// Otherwise, store the current number and its index
		numToIndex[num] = i
	}

	return []int{}
}
