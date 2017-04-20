package main

import (
	"fmt"
	"container/list"
)

func main() {
	var x list.List
	x.PushBack("Primeiro")
	x.PushBack("Segundo")
	x.PushBack("Terceiro")

	for e := x.Front(); e != nil; e =e.Next() {
		fmt.Println(e.Value.(string))
	}
}
