package main

import "fmt"

func main() {
	rst := f(10)
	fmt.Println(rst)
}

func f(x int) int {
	if x == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}
	
	return f(x-1)+f(x-2)

}

//함수 호출 과정
//1. 인자 기록
//2. IP 를 함수 시작점으로 이동
//3. return시 처음 IP로 돌아감.

