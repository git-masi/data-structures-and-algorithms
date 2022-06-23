// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/

// This was harder than expected. I didn't get a working solution for a few
// iterations of the code. Initially I thought it would be possible to sort
// the slice and use the index and value at the index to determine the answer
// and that still might be possible but I couldn't get it to work.

// The current solution uses O(n^2) time and O(n) space

// Looking at the discussion it seems like sorting the array is the way to go.
// I'll need to try again at another time.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(specialArray([]int{3, 5}))
	fmt.Println(specialArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(specialArray([]int{0, 0}))
	fmt.Println(specialArray([]int{0, 4, 3, 0, 4}))
	fmt.Println(specialArray([]int{3, 6, 7, 7, 0}))
}

func specialArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	freq := make([]int, len(nums))

	for _, v := range nums {
		for i := 0; i < v && i < len(freq); i++ {
			freq[i] += 1
		}
	}

	for i, v := range freq {
		if i+1 == v {
			return v
		}
	}

	return -1
}
