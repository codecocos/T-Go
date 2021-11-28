package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}


func (s *Student) PrintSungjuk() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputSungjuk(class string,grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var s Student = Student{name:"Tucker", age:23, class: "수학", grade: "A+"}

	s.InputSungjuk("과학","C") 
	s.PrintSungjuk() 
	// (s Student) 수학 A+ : 과학 C 로 값이 변경되지 않음.
	// (s *Student) 과학 C : 원하는 값으로 변경 됨.

}

//c,c++
// . : s.grade의 값(value)을 변경(value 형태 일 때만 . 사용)
// -> :  p -> grade = "c" (메모리주소가 가리키는 곳의 속성을 변경하고 싶을 때 화살표 사용)

//go
// . 만 사용, 컴파일러가 알아서 구분해줌
// 포인터타입이 . 을 사용하면 메모리주소가 가르키고 있는 곳의 속성을 변경,
// 그냥 value 타입이면 그냥 그 value의 값을 변경시킴.


//go 에서
//포인터를 함수 인자로 받으면, 메모리 주소만 복사 됨.
//값을 함수 인자로 받으면, 전체 속성이 복사 됨.