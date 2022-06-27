// https://leetcode.com/problems/pascals-triangle/

// This was pretty easy overall. The difficult part was
// determining whether or not this was the best solution.
// The factor that lead me to think that this was the best
// solution was the fact that the question specifies
// returning all of the rows. So I needed to create all the
// rows which means that the space is already going to
// proportional to the input and as such it makes sense that
// the time would be as well.

// If the question asked to calculate the value of the nth
// row or to calculate the value of all rows up to n there
// might be a better solution. Perhaps one rooted in math.
// But that wasn't the ask so I didn't spend much time thinking
// about it.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(2))
	fmt.Println(generate(1))
}

func generate(numRows int) [][]int {
	result := [][]int{{1}}

	if numRows == 1 {
		return result
	}

	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)

		row[0] = 1
		row[len(row)-1] = 1

		for j := 1; j < len(row)-1; j++ {
			row[j] = result[i-1][j] + result[i-1][j-1]
		}

		result = append(result, row)
	}

	return result
}
