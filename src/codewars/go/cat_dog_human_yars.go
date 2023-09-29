package main

import (
	"fmt"
)

func main() {
	test := 2
	fmt.Println(CalculateYears(test))
}

func CalculateYears(years int) (result [3]int) {
	result[0] = years
	for i := 1; i <= years; i++ {
		if i == 1 {
			result[1] += 15
			result[2] += 15
		} else if i == 2 {
			result[1] += 9
			result[2] += 9
		} else {
			result[1] += 4
			result[2] += 5
		}

	}

	return result
}
