package main

import (
	"fmt"

	"github.com/T-Go/dataStruct"
)

func main() {
	list := &dataStruct.LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		//list struct 안에 root와 tail이 있음.
		list.AddNode(i)
	}
	
	list.PrintNodes()
	
	//Root 다음을 지움
	list.RemoveNode(list.Root.Next)	
	list.PrintNodes()

	//루트를 지움
	list.RemoveNode(list.Root)
	list.PrintNodes()

	//테일을 지움
	list.RemoveNode(list.Tail)
	list.PrintNodes()
	fmt.Printf("Tail:%d\n", list.Tail.Val)

	list.PrintReverse()

	//slice////////////////////////////////////////////////////////////////////
		// remove 가 맨 앞과 맨 끝에서 일어 날 떄는 O(1), 중간에 일어나면 O(N)
		// 
		a := []int{1,2,3,4,5}
		//a[0:2] : 인덱스 0 부터 2번째 까지 니까 1과 2
		//a[3:]... : 인덱스 2부터 쭉.
		a = append(a[0:2], a[3:]...)
		fmt.Println(a)
	//////////////////////////////////////////////////////////////////////////

	//double linked list///////////////////////////////////////////////////////


	//특정한 인덱스의 값을 가져오고 싶다
	//슬라이스는 곧바로 가져올 수 있음. 
	//시작 메모리주소를 알고, 각 항목의 사이즈를 앎. 인덱스*SIZE+시작주소 : 이 주소로 곧바로 가면됨. O(1)
	//링크드리스트는 원하는 값이 나올 때 까지 전진: O(N)

	//Random Access
	//slice : O(1)
	//list : O(N)

	//배열은 캐시(4kb)(근방의 메모리)를 활용하는 장점이 있음.
	//리스트는 뚝뚝 떨어져 있는 메모리를 연결 시켜놓음- 캐시미스의 확률이 높음.
	

}
