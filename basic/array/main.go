package main

import "fmt"

func main() {
	// // var 배열이름 [배열크기]자료형
	// // 배열의 크기는 자료형을 구성하는 한 요소
	// var A [10]int

	// for i := 0; i < len(A); i++ {
	// 	A[i] = i * i
	// }

	// fmt.Println(A)
	
	// // [...]을 이용한 배열 크기 자동 설정
	// var arr = [...]int{9, 8, 7, 6} 
	// //arr 전체와  배열 크기 출력해보기
	// fmt.Println(arr, len(arr))    

	/////////////////////////////////////////////

	s := "Hello 월드"

	fmt.Println("len(s) = ", len(s))
	for i := 0; i < len(s); i++ {
		//fmt.Print(s[i],",")
		fmt.Print(string(s[i]),",")
	}

	fmt.Println("")

	s2 := []rune(s)
	fmt.Println("len(s2) = ", len(s2))
	for i := 0; i < len(s2); i++ {
		//fmt.Print(s2[i],",")
		fmt.Print(string(s2[i]),",")
	}
}

//len(A) : 배열의 길이

//문자열도 배열
//byte == uint8
//문자 하나는 1~3byte
//0~255

//배열 : 메모리
//길이 = 항목길이*갯수

// []rune : 고랭의 변수 타입, UTF-8의 문자를 나타냄