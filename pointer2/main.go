package main

import "fmt"

func main() {
	var a int
	a = 1

	//Increase 함수의 x와 a 는 서로 다른 변수(서로 다른 메모리공간을 사용)
	Increase(a)
	
	fmt.Println(a) //1

	
	Increase2(&a)

	fmt.Println(a) // 2
}


func Increase(x int) {
	x++ 
}


//x가 가지는 값은 메모리주소
func Increase2(x *int)  {
	//x++
	//x* : x의 메모리주소가 가르키는 곳의 값
	*x = *x + 1
}