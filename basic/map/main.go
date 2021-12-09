package main

import "fmt"

// type Key struct {
// 	v int
// }

// type Value struct {
// 	v int
// }

func main() {
	//[키타입]밸류타입
	var m map[string]string
	//var m1 map[int]string
	// var m2 map[Key]Value
	// var m3 map[Key]*Value

	//바로사용 불가능 : m이 nil이다.
	//초기화가 필요함.
	m= make(map[string]string)
	//m[키] = 밸류
	m["abc"] = "bbb"

	fmt.Println(m["abc"])

	m1 := make(map[int]string)
	m1[53] = "ddd"
	fmt.Println(m1[53])
	
	//go에서 셋팅된 값이 없는 경우, 디폴트 값을 출력해준다.
	fmt.Println("m1[55] = ", m1[55])

	m2 := make(map[int]int)
	m2[4] = 4
	fmt.Println("m2[10] = ", m2[10])

	m2[5] = 0
	fmt.Println("m2[5] = ", m2[5])

	//ok 부분을 통해 설정된 값인지 아닌지 알 수 있다.
	v , ok1 := m2[5]
	//v1 := m2[4]
	v2, ok2 := m2[10] //설정한 적이 없으므로 ok 값이 false
	fmt.Println(v, ok1, v2, ok2)

	//(맵이름, 키값)
	delete(m2 ,5)
	v , ok1 = m2[5]
	fmt.Println(v, ok1, v2, ok2)

	m2[2] = 98
	m2[10] = 1245

	//순회하면서 출력
	//for range 구문은 여러개의 요소를 가진 배열이나 맵, 슬라이스 등을 각 요소를 순회할 때 사용.
	//키 값이 정렬된 형태로 반환 되는 것은 아님, 무작위.
	for key, value := range m2{
		fmt.Println(key, " = " , value)
	}

}