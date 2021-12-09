package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func (t *Student) SetName(newName string) {
	t.name = newName
}

//Instance : pointer 타입의 struct
//(t *Student)가 앞으로 가면서 주어의 역할 : Instance -> 생명을 가진다고 생각.
// func SetName(t *Student, name string)  {
// 	t.name = name
// }

func (t *Student) SetAge(age int) {
	t.age = age
}

func PrintStudent(u *Student)  {
	fmt.Println(u)
}
func main() {
	var a *Student
	a = &Student{"aaa", 20, 10}

	a.SetName("bbb")
	a.SetAge(30)
	
	PrintStudent(a)

}
