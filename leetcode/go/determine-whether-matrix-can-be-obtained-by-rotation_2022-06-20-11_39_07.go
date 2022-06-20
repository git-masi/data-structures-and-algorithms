// https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/

// Failed this attempt
// Conceptually I think my solution works but I am having a hard time writing the code for it

// The idea is that there are a number of "layers like an onion and we can use
// these layers to compare our input 2d array with the target

// Take this example:
// mat			target
// 0 1 1		1 1 0
// 1 0 1		1 0 1
// 1 1 0		0 1 1
//
// If we treat the outermost numbers of `target` as a "layer" we can lay them out as a 1D array:
//
// * ----> |
// | 1 1 0 |
// | 1 0 1 |
// | 0 1 1 |
// | <---- |
//
// 1 1 0 1 1 1 0 1
//
// Then we can start at each corner of `mat` and see if we can get the values to line up with `target`
// So from top left we try to compare:
// 0 1 1 1 0 1 1 1
//
// If come upon a value that isn't equivalent we move on to the next corner
// If none of the corners works then we return false, no need to check further
// If we find a corner where all values match up we move to the next layer
// but start from the same corner and don't check other corners
// It helps to visualize with a larger 2d array

// L1 ---------------------->
// 0 			1			1		 	0			1
// 				L2 --------->
// 0			1			1			1 		0
// 							L3
// 1			0			0			0			0
//
// 0			0			0			0			0
//
// 1			1			1			1			1

// Again, conceptually it works it's just hard to code
// I need more practice with matrix problems

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's write some bugs!")
}

func findRotation(mat [][]int, target [][]int) bool {
	top := 0
	right := len(mat[0]) - 1
	bottom := right // n x n 2d array so all sides are equal
	left := 0

	for {
		// need condition where we reach the end
		if top > bottom || right > left {
			break
		}

		for i, v := range mat[top] {
			if v != target[top][i] {
				break // placeholder
			}
		}

		top--

		for i := top + 1; i < bottom; i++ {
			if mat[i][right] != target[i][right] {
				break // placeholder
			}
		}

		right--

		for i, v := range mat[bottom] {
			if v != target[bottom][i] {
				break // placeholder
			}
		}

		bottom++

		for i := top + 1; i < bottom; i++ {
			if mat[i][left] != target[i][left] {
				break // placeholder
			}
		}

		left++
	}

	return false
}
