package main

import (
	"fmt"
)

func main() {
	test := []int{1, 2, 3}
	result := Maps(test)
	fmt.Println(result)
}

func Maps(x []int) []int {
	var output []int
	for _, v := range x {
		output = append(output, v*2)
	}
	return output
}
