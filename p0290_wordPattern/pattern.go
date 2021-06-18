package pattern

import "strings"

func wordPattern(pattern string, s string) bool {
	// two map approach
	// pm tracks the individual chars in pattern, and the index they were found
	pm := make(map[rune]int)
	// sm tracks the space separated strings in s, and the index they were found
	sm := make(map[string]int)

	// split s into separate inner strings, sep by commas
	ss := strings.Split(s, " ")

	if len(ss) != len(pattern) {
		return false
	}
	// convert pattern to runes to be able to access individual values of indeces in pattern
	r := []rune(pattern)

	for i, c := range r {
		// if c (char) in r (pattern) doesnt exist, add it to its map with the index as value
		if _, ok := pm[c]; !ok {
			pm[c] = i
		}
		// if individual word within s doesnt exist, add it to its map with index as value
		if _, ok := sm[ss[i]]; !ok {
			sm[ss[i]] = i
		}
		// now check if the index values at each map mistmatch, if so, then we may return false
		if pm[c] != sm[ss[i]] {
			return false
		}
	}
	return true
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Pattern.
//Memory Usage: 2 MB, less than 21.15% of Go online submissions for Word Pattern.

// O(n) time, O(M) where M is unique words in s, and O(1) space for pattern since we can only have 26 chars (
//all lower case)
