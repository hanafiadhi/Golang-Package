package main

import (
	"fmt"
	"strconv"
)

func main() {
	var data int = 50
	result := []string{}
	// switch {
	// case data >= 50:
	// 	fmt.Println("Max number 50")
	// default:
	// 	for i := 1; i <= data; i++ {
	// 		if i%2 != 0 {
	// 			result = append(result, "Odd")
	// 			continue
	// 		}
	// 		co := strconv.FormatInt(int64(i), 10)
	// 		result = append(result, co)
	// 	}
	// }
	fmt.Println(result)
	if data > 50 {
		fmt.Println("Max number 50")
	} else {
		for i := 1; i <= data; i++ {
			if i%2 != 0 {
				result = append(result, "Odd")
				continue
			}
			co := strconv.FormatInt(int64(i), 10)
			result = append(result, co)
		}
		fmt.Println(result)
	}

}
