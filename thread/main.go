package main

import (
	"fmt"
	"time"
)

func main() {
	//go : go thread로 돌려라
	go fun1()
	go fun1()
	for i := 0; i < 20; i++ {
		fmt.Println("main", i)
	}
	fmt.Scanln()
}

func fun1() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 *time.Microsecond)
		fmt.Println("fun1:",i)
	}
}