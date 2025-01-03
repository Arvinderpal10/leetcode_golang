// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.

// Example 1:

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
// Example 2:

// Input: height = [1,1]
// Output: 1

// Constraints:

// n == height.length
// 2 <= n <= 105
// 0 <= height[i] <= 104
package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	MaxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			length := math.Min(float64(height[i]), float64(height[j]))
			breadth := math.Abs(float64(i - j))
			area := int(length * breadth)
			if area > MaxArea {
				MaxArea = area
			}
		}
	}
	return MaxArea

}
func main() {
	heights1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	heights3 := []int{1, 1}
	heights4 := []int{5}
	heights2 := []int{}
	fmt.Println(maxArea(heights1))
	fmt.Println(maxArea(heights2))
	fmt.Println(maxArea(heights3))
	fmt.Println(maxArea(heights4))
}
