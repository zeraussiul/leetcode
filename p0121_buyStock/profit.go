package profit

import "math"

func maxProfit(prices []int) int {
	// constraints found on problem: 1 <= prices.length <= 10^5
	if len(prices) <= 1 {
		return 0
	}

	// initialize a min and max to their opposites.
	min := math.MaxInt32
	max := 0
	// v = prices[i] as we iterate
	for _, v := range prices {
		// if v is less than the current min, change min to v
		if v < min {
			min = v
			// else if v - curr min is MORE than the curr max, change max
		} else if v-min > max {
			max = v - min
		}
	}
	return max
}

/*
7, 1, 5, 3, 6, 4

*/
