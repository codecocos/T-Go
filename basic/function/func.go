package main

import "fmt"

func main() {
	//1.
	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}
	
	//2.
	a,b := func1(2,3)
	fmt.Println(a,b)

	//3.
	f1(10)
	
	//4. 1부터 10까지의 합계
	s := sum(10,0)
	fmt.Println(s)

	//5. 재귀호출을 반복문으로 바꾸기
	s1 := 0
	for i := 0; i <= 10; i++ {
		s1 += i
	}
	fmt.Println("s1 : ", s1)

}

//1.함수
func add(x int, y int) int {
	return x + y
}

//2.함수안에서 다른 함수 호출 가능
//func1과 func2의 x,y는 별개b
func func1(x,y int) (int,int) {
	func2(x,y)
	return y,x

}

func func2(x,y int)  {
	fmt.Println("func2",x,y)
}

//3.재귀호출
func f1(x int)  {
	if x == 0 {
		return
	}
	fmt.Printf("f1(%d) before call f1(%d)\n",x, x-1)
	f1(x-1)
	fmt.Printf("f1(%d)after call f1(%d)\n",x, x-1)
}

//4.합계 구하기
func sum(x int, s int) int {
	if x == 0 {
		return s
	}
	// s = s+x
	s += x
	return sum(x-1,s)
}