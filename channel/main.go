package main

import "fmt"

func pop(c chan int)  {
	fmt.Println("Pop func")
	//채널에서 값을 하나 뺌.
	v := <-c
	fmt.Println(v)
}

func main() {
	//channel 기본
	//var c chan int

	//초기화 (chan 타입, 사이즈) : 사이즈 없는 경우 default는 0
	//c = make(chan int, 1)

	// // c 에 push
	// c <- 10
	// // c 에서 pop
	// v := <-c

	// fmt.Println(v)

	///////////////////////////////////////////////////////////////

	// //값을 넣을 때, 다른 쪽에서 뺴주지 않으면, 넣는 구문이 끝나지 않음.
	// //deadlock 발생
	// c1 := make(chan int)
	// c1 <- 10
	// v1 := <- c1
	
	// fmt.Println(v1)
	
	///////////////////////////////////////////////////////////////

	var c chan int 
	//사이즈는 기본값 0
	c = make(chan int)
	
	//pop함수를 go thread로 실행
	go pop(c)
	c <-10
	
	fmt.Println("end of program")
	
	
	///////////////////////////////////////////////////////////////
}

/*
go thread 1 - main()
go thread 2 - pop()

1. main()에서 pop() 실행 
2. go thread2
- pop은 빼는 것인데 뺄 것이 없는 경우, 대기상태가 됨.
	즉, 뺄 값이 생길 때 깨지 멈춰 있음.
3. go thread1에서 c <- 10을 push
4. go thread2에서 10을 출력
5. go thread1에서 end of program 출력


*/