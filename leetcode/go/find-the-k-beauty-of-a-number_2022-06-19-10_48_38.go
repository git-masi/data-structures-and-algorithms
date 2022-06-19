package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/find-the-k-beauty-of-a-number/

// The string based solution was obvious to me but it also became
// apparent early on that I didn't need to convert to a string at
// all. Instead I could just make heavy use of the % operator.
// This solution was interesting so I pursued it however it is
// less readable that working with strings IMO and the max int32
// string representation is 10 bytes so it doesn't save much in
// terms of memory. Still it works.

func main() {
	// fmt.Println(divisorSubstrings(240, 2))
	fmt.Println(divisorSubstrings(430043, 2))
}

func divisorSubstrings(num int, k int) int {
	if k == 0 {
		return 0
	}

	m := int(math.Pow10(k))

	if num%m == num {
		if num%int(math.Pow10(k-1)) == num {
			return 0
		}

		return 1
	}

	count := 0
	curr := num

	for {
		if curr%m != 0 && num%(curr%m) == 0 {
			count++
		}

		if m > curr {
			break
		}

		curr = curr / 10
	}

	return count
}
