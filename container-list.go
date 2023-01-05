package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Hanafi")
	data.PushBack("Adhi")
	data.PushBack("Prasetyo")

	/////to forward
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	//reverse

	for e:=data.Back();e !=nil; e= e.Prev(){
		fmt.Println(e.Value)
	}
}