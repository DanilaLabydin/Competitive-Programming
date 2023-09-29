package main

import (
	"fmt"
)

func main() {
	length := 2.0
	width := 7.0
	height := 5.0
	fmt.Println(GetVolume(length, height, width))
}

func GetVolume(length, width, height float64) float64 {
	return length * width * height
}
