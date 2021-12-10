package main

import (
	"fmt"

	"github.com/T-Go/basic/dataStruct"
)

func main() {
	stack := []int{}

	//스텍 슬라이스에 추가함 
	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}
	
	fmt.Println(stack)

	for len(stack) > 0 {
		var last int

		//stack[startPoint:endPoint]
		//stack[:len(stack)-1] : 스텍은 맨 처음부터 시작해서
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}

	queue := []int{}

	//큐 슬라이스에 추가함 
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}

	fmt.Println(queue)

	for len(queue) > 0 {
		var front int
		//stack[:len(stack)-1] : 스텍은 맨 처음부터 시작해서
		front, queue = queue[0], queue[1:]
		fmt.Println(front)
	}

	stack2 := dataStruct.NewStack()

	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}
	
	fmt.Println("NewStack")

	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d ->", val)
	}

	queue2 := dataStruct.NewQueue()
	for i := 1; i <= 5; i++ {
		queue2.Push(i)
	}

	fmt.Println("NewQueue")

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d ->", val)
	}

} 