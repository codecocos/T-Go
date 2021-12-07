package main

import "fmt"

//인터페이스 : 외부 메소드(관계)가 어떤 것이 있는지 정의
// 즉, 관계를 따로 정의한 type
type SpoonOfJam interface {
	//이름(입력) 출력
	String() string
}

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

//Putjam(jam Jam) : Jam 메소드를 가진 모든 객체가 인자로 들어올 수 있음.
func (b *Bread) Putjam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread" + b.val
}

type StrawberryJam struct {
}

//GetOneSpoon() SpoonOfJam : Jam의 인터페이스를 포함하고 있다, 구현하고 있다.(implemented)
func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

type OrangeJam struct {
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

type AppleJam struct {

}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}


type SpoonOfStrawberryJam struct {
}

//String()을 가지고 있으므로 SpoonOfJam interface 를 이미 가지고 있음
func (s *SpoonOfStrawberryJam) String() string {
	return " + strawberry"
}

type SpoonOfOrangeJam struct{

}

func (s *SpoonOfOrangeJam) String() string {
	return " + Orange"
}

type SpoonOfAppleJam struct{

}

func (s *SpoonOfAppleJam) String() string {
	return " + Apple"
}


func main() {
	bread := &Bread{}
	//jam := &StrawberryJam{}
	//jam := &OrangeJam{}
	jam := &AppleJam{}
	bread.Putjam(jam)

	fmt.Println(bread)
}