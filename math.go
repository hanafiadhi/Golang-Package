package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))
	fmt.Println(math.Floor(1.3))
	fmt.Println(math.Ceil(1.3))
	fmt.Println(math.Max(13,33))
	fmt.Println(math.Min(13,33))
}