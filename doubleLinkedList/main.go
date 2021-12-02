package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
}

type LinkedList struct{
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	//맨 처음에 노드가 없을 때
	if l.root == nil {
		//노드를 하나 만들어서 노드의 주소를 루트에 넣어줌.
		l.root = &Node{val:val}
		l.tail = l.root
		return
	}
	//기존 테일의 뒤에 새로운 노드를 추가하고, 노드의 테일 값을 새로운 노드에 부여
	l.tail.next = &Node{val:val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root{
		l.root = l.root.next
		l.root.prev = nil
		node.next = nil
		return
	}

	//지우고자 하는 노드의 이전 노드로 가면됨.
	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		node.prev = nil
		prev.next = prev.next.next
		//새로운 넥스트의 prev를 현재 prev로 바꿔줘야 함.
		prev.next.prev = prev
	}

	node.next = nil
}

func (l *LinkedList) PrintNodes(){
	node := l.root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) PrintReverse() {
	node := l.tail
	
	for node.prev != nil {
		fmt.Printf("%d ->", node.val)
		node = node.prev
	}
	fmt.Printf("%d\n",node.val)
	
}

func main() {
	list := &LinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		//list struct 안에 root와 tail이 있음.
		list.AddNode(i)
	}
	
	list.PrintNodes()
	
	//root 다음을 지움
	list.RemoveNode(list.root.next)	
	list.PrintNodes()

	//루트를 지움
	list.RemoveNode(list.root)
	list.PrintNodes()

	//테일을 지움
	list.RemoveNode(list.tail)
	list.PrintNodes()
	fmt.Printf("tail:%d\n", list.tail.val)

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
