// Write a function that accepts an integer n and a string s as parameters, and returns a string of s repeated exactly n times.

// Examples (input -> output)
// 6, "I"     -> "IIIIII"
// 5, "Hello" -> "HelloHelloHelloHelloHello"
package main

import (
	"fmt"
)

func main() {
	str := "I"
	times := 5
	result := RepeatStr(times, str)
	fmt.Println(result)
}

func RepeatStr(repetitions int, value string) string {
	var result string
	for i := 0; i < 5; i++ {
		result += value
	}
	return result

}
