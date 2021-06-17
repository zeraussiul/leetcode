package longest

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	// map to track if we've visited a char before
	seen := make(map[byte]struct{})

	// vars
	count, left, right := 0, 0, 0

	for right < len(s) && left <= right {
		_, ok := seen[s[right]]
		// if we havent seen, we add to map, calc count, and increase right
		if !ok {
			//we havent seen, add to map
			seen[s[right]] = struct{}{}
			sum := right - left + 1
			if count < sum {
				count = sum
			}
			right++
		} else {
			// else we HAVE seen it, therefore we need to work from left side, delete left and increase
			delete(seen, s[left])
			left++
		}
	}
	return count
}
