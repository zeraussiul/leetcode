package maxunits

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	// fix as many units as possible.
	sort.Slice(boxTypes, func(i, j int) bool {
		// largest first
		return boxTypes[i][1] > boxTypes[j][1]
	})
	sum := 0
	for _, b := range boxTypes {
		if truckSize >= b[0] {
			// sum equals sum + box size * number of boxes of that size
			sum += b[0] * b[1]
			// subtract from trucksize the box size we just used
			truckSize -= b[0]
		} else {
			// we only have space for 'truckSize' boxes of size whatever we're on
			sum += truckSize * b[1]
			break
		}
	}
	return sum
}
