// https://leetcode.com/problems/valid-palindrome/

// This was relatively easy but I needed to look up how regular expressions
// work in Go. Also I think the code could be simplified if I convert the
// input string to lower case first.

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// fmt.Println(isPalindrome("racecar"))
	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("a."))
	fmt.Println(isPalindrome(" "))
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	start := 0
	end := len(s) - 1
	re := regexp.MustCompile(`[^A-Za-z0-9]`)

	for {
		for {
			if start < end && re.MatchString(s[end:end+1]) {
				end--
			} else {
				break
			}
		}

		for {
			if start < end && re.MatchString(s[start:start+1]) {
				start++
			} else {
				break
			}
		}

		if start > end {
			break
		}

		// Another option is:
		// 		unicode.ToLower(rune(s[start])) != unicode.ToLower(rune(s[end]))
		if strings.ToLower(s[start:start+1]) != strings.ToLower(s[end:end+1]) {
			return false
		}

		end--
		start++
	}

	return true
}
