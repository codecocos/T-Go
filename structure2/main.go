package main

import "fmt"

//Student type 생성
type Student struct {
	name  string
	class int

	sungjuk Sungjuk
}

type Sungjuk struct {
	name  string
	grade string
}

//메서드가 구조체 바깥에서 정의 됨.
//구조체(객체)에 속한 메서드
func (s Student) ViewSungjuk() {
	fmt.Println(s.sungjuk)
}


//일반함수 : 객체에 속하지 않은 함수
func ViewSungjuk(s Student)  {
	fmt.Println(s.sungjuk)
}

func (s Student) InputSungjuk (name string, grade string){
	s.sungjuk.name = name
	s.sungjuk.grade = grade
	
}

func InputSungjuk(s Student, name string, grade string)  {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

func main() {

	var s Student
	s.name = "coco"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "A"

	//메서드
	s.ViewSungjuk()

	//일반함수
	ViewSungjuk(s)

	//여기 s 와 30번째 줄의 s와는 다름. 메모리주소가 서로 다름. 값만 복사되어 같을 뿐이다. 
	//31번의 s 의 값을 바꿨다고 해도, 58번의 s 값이 바뀌는 것은 아니다. 
	s.InputSungjuk("과학","B")
	s.ViewSungjuk() // {수학 A} ***주의 : {과학 B}가 아님.


}

//고랭에서 함수 호출 : 변수는 무조건 복사로!
//복사가 이뤄질 때는 값이 복사되는 것이지, 그 메모리가 그대로 전달 되는 것은 아님.
