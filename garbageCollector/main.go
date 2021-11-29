package main

import "fmt"

func add() *int {
	//garbage collector : 불필요한 메모리를 정리해 줌.
	//메모리가 정리되는 기준점 : 메모리(변수)가 필요 없어지는 순간
	//즉, 'reference count = 0' 이 되는 순간 GC가 지움.

	//reference count
	var a int //1
	var p *int

	a = 3
	p = &a //2

	fmt.Println(*p)

	//1.리턴되는 순간 메모리가 사라지는 것이아니라,
	return p
}

func main() {

	//2. 실행이 끝나는 순간 메모리가 정리됨.
	v := add()

	fmt.Println(v)
}

// GC가 없는 경우(unmanaged language) : C , C++

// C언어
// 스택 메모리 : 변수의 생명주기에 맞는 스코프 범위를 벗어나면 무조건 사라져야 함.
// 힙 메모리 : 프로그래머가 힙 메모리 할당 가능, 할당방법은 malloc(사이즈)이용, 그리고 사용하고 나서는 free로 메모리 삭제

// 버그들
// memory leak: 흔히들 하는 실수.
// dangiling : 사라진 메모리를 참조하는 경우

//reference count가 0 이 아님에도 아무도 쓰지 않는 경우가 있음 : 외딴섬
// a->b->c->a 서로를 카운터하지만, 외부에서는 아무도 a,b,c를 가르키는 것이 없다.
// 위 의 경우도 메모리 leak, 가비지컬렉터가 지워줌.

// GC : 전수조사한다, 하는 일이 많다. 메모리를 많이 사용한다. 속도가 느리다. 
// 멀티쓰레드 발전으로 GC의 성능이 좋아짐.
// GC가 있어도 memory leak 이 발생할 수 있다.



 