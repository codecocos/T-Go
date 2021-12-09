package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

//함수 형태 : 기능 단위
//*Student 해주고 &a로 호출 하면 원본의 값을 변경가능.
func SetName(t *Student, newName string) {
	t.name = newName
}

//method로 변경 : object 단위
func (t *Student) SetName(newName string) {
	t.name = newName
}

/* 함수와 메소드는 같음, 다만 어느 object에 속했는 지 알려주는 게 메소드라고 생각하면 됨. */


func main() {
	a := Student{"aaa", 20, 10}

	// var b *Student
	// b = &a

	// //b가 포인팅하고 있는 값을 바꾸겠다.
	// //reference type : pointing하고 있는 형태가 사용됨
	// b.age = 30

	// fmt.Println(a)
	// fmt.Println(b)
	
	//함수호출
	//SetName(&a, "bbb") 

	//메소드 호출
	a.SetName("bbb") //subject.verb(object) -> subject를 Instance(생명주기)라고 함.

	fmt.Println(a) //(t Student)일 때 {aaa 20 10} : a의 이름이 바뀌지 않음 -> 서로 다른 메모리 공간
}

///////////////////////////////////////////////////////////////////////////////////////////////////

//value type : 값 형태로 넘기는 것 
//->프로퍼티를 복사 : 서로 다른 메모리 공간

//reference type : 포인터로 넘기는 것
// -> 메모리 주소를 복사 : 같은 메모리를 가르킴