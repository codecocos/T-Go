package dataStruct

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

//첫글자가 대문자 : 외부 공개
//첫글가사 소문자 : 외부 공개 안함.
type LinkedList struct{
	Root *Node
	Tail *Node
}

func (l *LinkedList) AddNode(val int) {
	//맨 처음에 노드가 없을 때
	if l.Root == nil {
		//노드를 하나 만들어서 노드의 주소를 루트에 넣어줌.
		l.Root = &Node{Val:val}
		l.Tail = l.Root
		return
	}
	//기존 테일의 뒤에 새로운 노드를 추가하고, 노드의 테일 값을 새로운 노드에 부여
	l.Tail.Next = &Node{Val:val}
	Prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = Prev
}

func (l *LinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val
	}
	return 0
}

func (l *LinkedList) Front() int {
	if l.Root != nil {
		return l.Root.Val
	}
	return 0
}

func (l *LinkedList) Empty() bool {
	return l.Root == nil
}

func (l *LinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}

func (l *LinkedList) PopFront() {
	if l.Root == nil {
		return
	}
	l.RemoveNode(l.Root)
}



func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root{
		l.Root = l.Root.Next
		if l.Root != nil {
			l.Root.Prev = nil
		}
		node.Next = nil
		return
	}

	//지우고자 하는 노드의 이전 노드로 가면됨.
	Prev := node.Prev

	if node == l.Tail {
		Prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = Prev
	} else {
		node.Prev = nil
		Prev.Next = Prev.Next.Next
		//새로운 넥스트의 prev를 현재 prev로 바꿔줘야 함.
		Prev.Next.Prev = Prev
	}

	node.Next = nil
}

func (l *LinkedList) PrintNodes(){
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

func (l *LinkedList) PrintReverse() {
	node := l.Tail
	
	for node.Prev != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Prev
	}
	fmt.Printf("%d\n",node.Val)
	
}