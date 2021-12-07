package main

import "fmt"

type Bread struct {
	val string
}

type Jam struct {
}

func (b *Bread) PutJam(jam *Jam) {
	b.val += jam.GetVal()
}

func (b *Bread) String() string {
	return b.val
}

func (j *Jam) GetVal() string {
	return " + jam"
}

func main() {
	//브래드 인스턴스 생성
	bread := &Bread{val: "bread"}
	//잼 인스턴스 생성
	jam := &Jam{}

	//브래드 인스턴스에 속해 있는 함수인 PutJam 메소드 호출
	bread.PutJam(jam)

	//인자로 넘어간 인스턴스에 스트링이란 메소드가 있을 경우,
	// 스트링 메소드의 결과값을 출력
	fmt.Println(bread)
}