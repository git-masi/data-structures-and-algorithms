//https://leetcode.com/problems/largest-3-same-digit-number-in-string/

// This was quite easy. Standard sliding window problem.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestGoodInteger("6777133339"))
}

func largestGoodInteger(num string) string {
	if len(num) < 3 {
		return ""
	}

	start := 0
	end := 1
	result := ""

	for {
		if end >= len(num) {
			break
		}

		if num[start] == num[end] {
			end++
		} else {
			start = end
			end++
		}

		if end-start == 3 {
			if result == "" || int(num[start]) > int(result[0]) {
				result = num[start:end]
			}

			start = end
			end++
		}
	}

	return result
}
