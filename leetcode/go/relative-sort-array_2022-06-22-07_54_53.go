// https://leetcode.com/problems/relative-sort-array/

// This was a little harder than expected. Partly because I didn't
// know the sytax for Go sorting or copying a portion of an array.
// Also a more optimal solution using a hashmap didn't occur to me.
// The current solution is O(n * m + (n-m)log(n-m)) so not ideal.

package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
	fmt.Println(relativeSortArray([]int{33, 22, 48, 4, 39, 36, 41, 47, 15, 45}, []int{22, 33, 48, 4}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	k := 0

	for i := 0; i < len(arr2); i++ {
		for j := k; j < len(arr1); j++ {
			if arr1[j] == arr2[i] {
				arr1[k], arr1[j] = arr1[j], arr1[k]
				k++
			}
		}
	}

	temp := make([]int, len(arr1)-k)

	copy(temp, arr1[k:])

	// This could be replaced with `sort.Ints(temp)`
	sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })

	copy(arr1[k:], temp)

	return arr1
}
