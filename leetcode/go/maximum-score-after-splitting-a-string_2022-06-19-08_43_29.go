package main

import (
	"fmt"
)

// https://leetcode.com/problems/maximum-score-after-splitting-a-string/

// This solution was relatively easy for me to come up with
// but given that it requires two iterations I wasn't sure if it was
// the optimal solution so I had to spend time thinking of other
// possibilities.
// I had considered sliding windows and fast/slow pointer solutions
// but I couldn't come up with anything faster. However, given that
// this is asymptotically equal to O(n) time and space is O(1) I
// ultimately concluded that it was a good solution.
// I haven't seen anything on leetcode that suggests something faster
// as of right now so I think it's optimal.

func main() {
	s := "0001000011"

	fmt.Println(maxScore(s))

	s = "011101"

	fmt.Println(maxScore(s))
}

func maxScore(s string) int {
	// Handle edge case where the string has less than 2 characters
	if len(s) < 2 {
		return 0
	}

	max := 0
	i := 0

	if s[i] == '0' {
		max++
	}

	for i = 1; i < len(s); i++ {
		if s[i] == '1' {
			max++
		}
	}

	curr := max

	for i = 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			curr++
		}

		if s[i] == '1' {
			curr--
		}

		if curr > max {
			max = curr
		}
	}

	return max
}
