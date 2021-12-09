package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	//---------------------------------------------------------------------------
	
	
	//1. 배열의 복사
	clone := [5]int{}
	
	for i := 0; i < 5; i++ {
		clone[i] = arr[i]
	}

	fmt.Println(clone)

	//////////////////////////////////////////////////////

	//2. 배열의 역순
	//2-1. 복사 후 대입(2N번 복사)
	temp := [5]int{}
	
	for j := 0; j < 5; j++ {
		//역순으로 복사 후
		temp[j] = arr[len(arr)-1-j]
	}

	for k := 0; k < len(arr); k++ {
		//복사된 값을 대입
		arr[k] = temp[k]
	}

	fmt.Println("temp:",temp)
	fmt.Println("arr:",arr)

	//2-2. 값을 교환 후 대입(N/2 복사)

	arr2 := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr2)/2; i++ {
		//이중배열
		arr2[i], arr2[len(arr)-1-i] = arr2[len(arr)-1-i], arr2[i]
	}

	fmt.Println(arr2)

}