package main

import "fmt"

func main() {
	//slice (dynamic Array) : structure value
	var s []int
	var t []int

	s = make([]int, 3)

	s[0] = 100
	s[1] = 200
	s[2] = 300

	//첫번째 나오는 슬라이스에, 뒤에 것들을 계속 붙임
	//s = append(s, 400, 500, 600, 700)
	t = append(s, 400)

	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))

	fmt.Println("///////////////////////")
	var u []int
	
	u = append(t, 500)
	
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))//같은 메모리 주소
	fmt.Println(u, len(u), cap(u))//같은 메모리 주소

	u[0] = 9999

	fmt.Println("///////////////////////")
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
	fmt.Println(u, len(u), cap(u))

}

// type Slice struct{
// 	pointer * //시작주소
// 	len int		//갯수
// 	cap int		//할당된 최대 갯수
// }

//시작주소에서 len의 수만큼 내 것인데 , 최대 cap까지 사용가능하다.