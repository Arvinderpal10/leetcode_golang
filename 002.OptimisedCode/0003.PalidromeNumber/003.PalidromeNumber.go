// Given an integer x, return true if x is a palindrome, and false otherwise.

// Example 1:

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

// Constraints: -2^31 <= x <= 2^31 - 1

// Follow up: Could you solve it without converting the integer to a string?
package main

import "fmt"

func isPalindrome(x int) bool {
	numCopy := x
	sum := 0
	for x > 0 {
		digit := x % 10
		sum = sum*10 + digit
		x = x / 10
	}
	if numCopy == sum {
		return true
	} else {
		return false
	}
}

func main() {
	// 123 , 333, 454, 1212121
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(333))
	fmt.Println(isPalindrome(454))
	fmt.Println(isPalindrome(1212121))

}
