package main

import "fmt"

type Bread struct {
	val string
}

type StrawberryJam struct {
}

type SpoonOfStrawberryJam struct {
}

func (b *Bread) Putjam(jam *StrawberryJam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread" + b.val
}

func (j *StrawberryJam) GetOneSpoon() *SpoonOfStrawberryJam {
	return &SpoonOfStrawberryJam{}
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

func main() {
	bread := &Bread{}
	jam := &StrawberryJam{}
	bread.Putjam(jam)

	fmt.Println(bread)
}