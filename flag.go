package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")

	flag.Parse()

	fmt.Println(*host)
}
