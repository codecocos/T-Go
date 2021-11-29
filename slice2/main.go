package main

import "fmt"

func main() {
 //슬라이스에 append 시 메모리주소가 다른 경우와 같은 경우가
 //각각 있음으로 복사로 처음부터 메모리주소를 다르게 하자.

	//슬라이스 생성
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2

	//다른 슬라이스 생성 : 서로 다른 메모리를 확보.
	b := make([]int, len(a))

	//a 슬라이스를 b슬라이스로 복사 
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}

	//b 슬라이스에 3을 추가
	b = append(b, 3)

	fmt.Printf("%p %p\n", a, b)

	b[0] = 4
	b[1] = 5

	//b의 값을 바꿔도 a값은 변경이 없다.
	fmt.Println(a)
	fmt.Println(b)
}