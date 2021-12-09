package main

//////////////////////////////////////////////////
// 1.

// type Actor struct {
// }

// func (a *Actor) Move() {
// }

// func (a *Actor) Attack() {
// }

// func (a *Actor) Talk() {
// }

//////////////////////////////////////////////////
// 2. 묶여 있으면 의존성을 발생시키기 쉽다.

type Actor interface {
	Move()
	Attack()
	Talk()
}



//////////////////////////////////////////////////
// 3.
type Talkable interface{
	Talk()
}

type Attackable interface{
  Attack()
}

type Movalbe interface{
	Move()
}

//srp에 어긋남.
func MoveTo(a Actor)  {
	a.Move()
	a.Attack()
}

//srp에 어긋나지 않도록 강제
func Moveto(a Movalbe)  {
	a.Move()
	//Movalbe 인터페이스에 정의되어 있지 않음.
	//a.Attack()
}
