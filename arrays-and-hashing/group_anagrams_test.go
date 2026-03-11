package arrays

import (
	"slices"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}

	tests := []struct {
		name     string
		args     args
		expected [][]string
	}{
		{
			name: "Example 1",
			args: args{
				strs: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			},
			expected: [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.args.strs)

			if !compare2D(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// compare2D safely compares two 2D string slices regardless of order
func compare2D(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, group := range a {
		slices.Sort(group)
	}
	for _, group := range b {
		slices.Sort(group)
	}

	// 2. Sort the outer groups so they are in a predictable order
	sortFunc := func(g1, g2 []string) int {
		return slices.Compare(g1, g2)
	}
	slices.SortFunc(a, sortFunc)
	slices.SortFunc(b, sortFunc)

	// 3. Now that everything is sorted, we can safely compare them
	return slices.EqualFunc(a, b, slices.Equal)
}

// Make a slice of letters for each word
// add the word to the map of lettersToWords
// iterate through the map, and return the grouped words
func groupAnagrams(strs []string) [][]string {
	lettersToWords := make(map[[26]int][]string)

	for _, word := range strs {
		var letters [26]int
		for _, letter := range word {
			letters[letter-'a']++
		}
		lettersToWords[letters] = append(lettersToWords[letters], word)
	}

	wordGroups := make([][]string, 0, len(lettersToWords))
	for _, v := range lettersToWords {
		wordGroups = append(wordGroups, v)
	}

	return wordGroups
}
