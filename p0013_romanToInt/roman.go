package roman

func romanToInt(s string) int {
	// make a map of values, uint8 bc s[n] is of type uint8
	r := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// if s of length 1, just return the value of that single char
	if len(s) == 1 {
		return r[s[0]]
	}
	// start at end
	sum := r[s[len(s)-1]]
	// iterate backwards, adding values from map UNLESS we encounter a value that is less than the previous visited
	//one, for example IV, we visit V first, but I is less than V therefore we sub, not add.
	for i := len(s) - 2; i >= 0; i-- {
		if r[s[i]] < r[s[i+1]] {
			sum -= r[s[i]]
		} else {
			sum += r[s[i]]
		}
	}
	return sum
}
