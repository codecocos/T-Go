package main

import (
	"fmt"

	"github.com/T-Go/basic/dataStruct"
)

func main() {
	h := &dataStruct.MinHeap{}

	//[-1, 3, -1, 5, 4], 2번째 큰 값 찾기
	nums := []int{-1, 3, -1, 5, 4}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 2{
			h.Pop()
		}
	}
	fmt.Println(h.Pop()) // 4

	//Input: [2, 4, -2, -3, 8], 1
	//Output : 8
	h = &dataStruct.MinHeap{}

	nums = []int{2, 4, -2, -3, 8}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 1{
			h.Pop()
		}
	}
	fmt.Println(h.Pop())

	//Input: [-5, -3, 1], 3
	//Output : -5
	h = &dataStruct.MinHeap{}

	nums = []int{-5, -3, 1}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 3{
			h.Pop()
		}
	}
	fmt.Println(h.Pop())
}