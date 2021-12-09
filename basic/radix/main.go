package main

import "fmt"

func main() {
	arr := [11]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4, 5}
	//임시배열 : 항상 배열의 갯수가 같은 것은 아님, 0~9 로 들어가는 원소의 범위와 상관있음.
	//배열의 초기 값이 지정되어 있지 않아 기본 값인 0이 초기 값.
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		fmt.Println(idx)
		temp[idx]++
	}

	//temp[0]=2 : 0이 2개 있다는 의미.
	//arr[0] = 0, arr[1] = 0

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println(arr)
}

/* 
Radix Sort

낮은 자리수부터 비교하여 정렬해 간다는 것을 기본 개념으로 하는 정렬 알고리즘

원소 값의 범위가 한정적 일 때,


*/