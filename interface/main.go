package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
}

func (a *StructA) AAA(x int) int {
	return x * x
}

func (a *StructA) BBB(x int) string {
	//숫자를 문자로 반환
	return "X=" + strconv.Itoa(x)
}

type StructB struct{

}

func (b *StructB) AAA(x int) int {
	return x*2
}

func main() {
	var c InterfaceA
	c = &StructA{}

	// //structB는 interfaceA의 구현체가 아님
	// var d interfaceA
	// d = &StructB{}

	fmt.Println(c.BBB(3))
}