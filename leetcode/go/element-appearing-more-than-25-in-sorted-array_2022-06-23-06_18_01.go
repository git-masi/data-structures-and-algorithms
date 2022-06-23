// https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/

// This question was easy but the edge cases are a little deceptive so it's
// important to remember not to get too cocky else everything could get throw
// off by a tiny bug/edge case.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
	fmt.Println(findSpecialInteger([]int{1, 2, 3, 3}))
}

func findSpecialInteger(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	t := len(arr)/4 + 1

	start := 0
	end := 0

	for {
		end++

		if end >= len(arr) {
			break
		}

		for {
			if arr[start] == arr[end] {
				break
			}

			start++
		}

		if end-start+1 >= t {
			return arr[start]
		}
	}

	// We should never reach this case
	panic("Invalid slice passed as input")
}
