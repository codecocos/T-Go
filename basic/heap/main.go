package main

import (
	"fmt"

	"github.com/T-Go/basic/dataStruct"
)

func main() {
	h := &dataStruct.Heap{}

	h.Push(2)
	h.Push(6)
	h.Push(9)
	h.Push(6)
	h.Push(7)
	h.Push(8)

	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
} 


/*

9
| \
7  8
|\  \
2 6  6

*/