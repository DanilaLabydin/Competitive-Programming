// Clock shows h hours, m minutes and s seconds after midnight.

// Your task is to write a function which returns the time since midnight in milliseconds.

// Example:
// h = 0
// m = 1
// s = 1

// result = 61000
// Input constraints:

// 0 <= h <= 23
// 0 <= m <= 59
// 0 <= s <= 59

package main

import (
	"fmt"
)

func main() {
	var hour int = 0
	var minute int = 1
	var second int = 1
	var result int
	result = Past(hour, minute, second)
	fmt.Println(result)
}

func Past(h int, m int, s int) int {
	var result int
	result += s * 1000
	result += m * 60 * 1000
	result += h * 60 * 60 * 1000
	return result
}
