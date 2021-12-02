package main

import "fmt"


type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node

	root = &Node{val: 0}

	for i := 1; i < 10; i++ {
		AddNode(root, i)
	}

	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func AddNode(root *Node, val int) {
	//맨 끝의 노드를 처음 부터 시작해서 넥스트 노드가 없을 때까지 전진 : O(N)
	var tail *Node
	tail = root
	for tail.next != nil {
		tail = tail.next
	}

	//새로운 노드를 만들에서 제일 마지막 노드에 추가
	node := &Node{val: val}
	tail.next = node
}