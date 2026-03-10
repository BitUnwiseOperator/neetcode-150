// Package arrays has all of the array and hash questions from neetcode
package arrays

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	// 1. Define the input arguments
	type args struct {
		str1 string
		str2 string
	}

	// 2. Define the test table
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			name:     "Example 1: Standard success case",
			args:     args{str1: "racecar", str2: "carrace"},
			expected: true,
		},
		{
			name:     "Example 2: Standard failure case",
			args:     args{str1: "jar", str2: "jam"},
			expected: false,
		},
	}

	// 3. Iterate over the table and run subtests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validAnagram(tt.args.str1, tt.args.str2)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// to solve this we should start with
// iterating through the two strings and seeing
// if the array left over has only 0 values
func validAnagram(str1, str2 string) bool {
	// define the memory to make an array so we
	// can remain on the stack. 26 represents all the chars in the alphabet
	var count [26]int
	// if they aren't the same size we know they arent the same
	if len(str1) != len(str2) {
		return false
	}

	// using the array, iterate through both strings and
	// add and subtract from their location in the array
	for i := 0; i < len(str1); i++ {
		count[str1[i]-'a']++
		count[str2[i]-'a']--
	}

	// check to see if the array has only 0s, if it does, its a valid anagram
	for _, value := range count {
		if value != 0 {
			return false
		}
	}
	return true
}
