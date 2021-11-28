package main

import "fmt"

func main() {
	var a int
	var b int

	var p *int

	p = &a
	a = 3 // 어딘가의 메모리에 3이 적힘.
	b = 2


	fmt.Println(a) //a의 값이 출력됨
	fmt.Println(p) //a의 메모리주소가 출력됨.
	fmt.Println(*p)//a의 메모리주소가 가르키고 있는 주소의 값을 알고 싶을 때 : *p

	p = &b

	fmt.Println(b) //b의 값이 출력됨
	fmt.Println(p) //b의 메모리주소가 출력됨.
	fmt.Println(*p)//b의 메모리주소가 가르키고 있는 주소의 값을 알고 싶을 때 : *p
}