package merge

import "sort"

func merge(intervals [][]int) [][]int {
	// base check
	if len(intervals) <= 1 {
		return intervals
	}
	// sort intervals on rows
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	results := [][]int{}
	curr := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > curr[1] {
			results = append(results, curr)
			curr = intervals[i]
		} else if intervals[i][1] > curr[1] {
			curr[1] = intervals[i][1]
		}
	}
	results = append(results, curr)
	return results
}

/*
curr: [1,4]
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]

*/
