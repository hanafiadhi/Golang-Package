package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)
	//contoh go run os.go Hanafi Adhi Prasetyo
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])
}
