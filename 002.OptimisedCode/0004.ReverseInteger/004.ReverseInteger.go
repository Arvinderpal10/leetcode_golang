// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:
// Input: x = 123
// Output: 321
// Example 2:

// Input: x = -123
// Output: -321
// Example 3:

// Input: x = 120
// Output: 21

// Constraints:

// -231 <= x <= 231 - 1
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	reversedNum := 0
	for x != 0 {
		digit := x % 10
		reversedNum = reversedNum*10 + digit
		x = x / 10
	}
	if reversedNum > math.MaxInt32 || reversedNum < math.MinInt32 {
		return 0
	} else {
		return reversedNum
	}

}

func main() {
	// 123, -123, 120, 1534236469, -2147483648
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(1200))
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-2147483648))

}
