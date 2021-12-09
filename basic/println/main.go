package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
	val string
}

func (s *StructA) String() string {
	return "Val:" + s.val
}

type StructB struct {
	val int
}

func (s *StructB) String() string {
	// s.val += 10
	// fmt.Println("StructB String()")
	return "StructB:" + strconv.Itoa(s.val)
}



type Printable interface {
	String() string
}

//Printable interface를 가지고 있는 객체만 들어 올 수 있다.
func Println(p Printable) {
	fmt.Println(p.String())
}

func main() {
	a := &StructA{val: "AAA"}

	//어떤 객체를 넘겼을 때, String()이라는 인터페이스가 없으면, 그냥 그 객체의 값을 출력
	//만약 String() 인터페이스가 있으면, String()객체의 리턴 값을 출력한다.
	fmt.Println(a)
	
	Println(a)

	b := &StructB{val: 10}
	Println(b)
}