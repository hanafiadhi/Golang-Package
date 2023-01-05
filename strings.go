package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hanafi","a"))
	fmt.Println(strings.ToLower("ADHI"))
	fmt.Println(strings.ToUpper("adhi"))
	fmt.Println(strings.Trim("          adhi prasetyo               ", " "))
	fmt.Println(strings.ReplaceAll( "ha ha ha ha ha","ha","ea"))

}