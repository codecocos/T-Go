package main

import "fmt"

type Person struct {
	name string
	age  int
}

//어떤 타입이 가지고 있는 기능인지 말함.
func (p Person) PrintName() {
	fmt.Print(p.name)
}

func main() {
	var p Person
	p1 := Person{"coco", 15}
	p2 := Person{name: "gogo", age: 21}
	p3 := Person{name: "꼬꼬"}
	p4 := Person{} //defalt값은 공백 과 0

	fmt.Println(p,p1,p2,p3,p4)

	p.name = "볼"
	p.age = 24

	fmt.Println(p)

	//p의 기능을 사용
	p.PrintName()

}