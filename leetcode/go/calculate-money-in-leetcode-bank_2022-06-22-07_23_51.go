// https://leetcode.com/problems/calculate-money-in-leetcode-bank/

// This was easy but it was basically a math problem. Without the n(n+1)/2
// formula the answer is not as easy to code.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalMoney(7))
	fmt.Println(totalMoney(14))
	fmt.Println(totalMoney(8))
	fmt.Println(totalMoney(2))
	fmt.Println(totalMoney(1337))
}

func totalMoney(n int) int {
	levels := n / 7
	mod := n % 7
	total := 0

	total += (28 * levels) + (7 * sumToN(levels-1))

	if mod != 0 {
		for i := 0; i < mod; i++ {
			total += levels + 1 + i
		}
	}

	return total
}

func sumToN(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return n * (n + 1) / 2
}
