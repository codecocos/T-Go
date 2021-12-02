package main

import "fmt"

type Node struct {
	next *Node
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
	l.tail = l.tail.next
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root{
		l.root = l.root.next
		node.next = nil
		return
	}

		prev := l.root
	//이전 노드의 next가 현재 노드가 아니면 찾을 떄까지 전진
	for prev.next != node{
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next = prev.next.next
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
}
