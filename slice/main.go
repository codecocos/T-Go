package main

import "fmt"

func main() {
	//slice : GO의 동적배열 - 길이가 점점 늘어나는 배열
	//길이와 공간

	//var a []int
	//a := []int{1,2,3,4,5}
	a := make([]int,0,8) //사용은 0, 확보된 공간은 8

	fmt.Printf("len(a) = %d\n", len(a)) //사용
	fmt.Printf("cap(a) = %d\n", cap(a)) //공간
	
	// a 슬라이스에 항목 1 을 추가 한, 슬라이스를 반환
	a = append(a, 1)
	
	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	///////////////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////

	//기존의 슬라이스와 값이 추가된 슬라이스 비교
	//주의점 : append 사용시 메모리 주소같은 주소가 반환 될 수 도 있고, 서로다른 주소가 반환될 수 도 있다.

	b := []int{1,2}

	//값을 추가한 슬라이스는 기존의 슬라이스와 다른 슬라이스일 수 있다.
	c := append(b,3)

	//서로 다른 주소 값이 나온다 : 서로 메모리 공간이 다르다.
	fmt.Printf("%p %p\n", b, c)

	for i := 0; i < len(b); i++ {
		fmt.Printf("%d,", b[i])
	}
	fmt.Println()
	
	for i := 0; i < len(c); i++ {
		fmt.Printf("%d,", c[i])
	}
	fmt.Println()
	
	fmt.Println(cap(b)," ",cap(c))

	//----------------------------------

	d := make([]int, 2, 4)
  d[0]=1
	d[1]=2

	e := append(d,3)

	//같은 메모리 주소
	fmt.Printf("%p %p\n", d, e)

	fmt.Println(d)
	fmt.Println(e)

	e[0]=4
	e[1]=5
	
	fmt.Println(d)
	fmt.Println(e)
}
