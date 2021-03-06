package main

import (
	"fmt"
	"time"
)

func push(c chan int)  {
	i := 0
	for{
		time.Sleep(1*time.Second)
		c <- i
		i++
	}
}

func main() {
	c := make(chan int)

	go push(c)

	timerChan := time.After(10*time.Second)
	tickTimerChan := time.Tick(2*time.Second)

	for {
		select {
			case v := <-c:
				fmt.Println(v)
			// default:
			// 	fmt.Println("Idle")
			// 	time.Sleep(1*time.Second)
			case <- timerChan:
				fmt.Println("Timeout")
				return
			case <-tickTimerChan:
				fmt.Println("Tick")
		}
	}
}
//time.tick : channel - 일정 시간 간격으로 주기적으로 반복
//time.After : channel - 특정시간이후 입력이 나오는 함수