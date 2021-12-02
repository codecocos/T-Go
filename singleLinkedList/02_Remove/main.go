package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	//노드가 추가 될 떄 마다 추가되는 노드로 테일을 이동 : O(1)
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}
	
	PrintNodes(root)
	
	//root 다음을 지움
	root, tail= RemoveNode(root.next, root, tail)	
	PrintNodes(root)

	//루트를 지움
	root, tail= RemoveNode(root, root, tail)
	PrintNodes(root)

	//테일을 지움
	root, tail= RemoveNode(tail, root, tail)
	PrintNodes(root)
	fmt.Printf("tail:%d\n",tail.val)
}

func AddNode(tail *Node, val int) *Node {

	//새로운 노드를 만들에서 제일 마지막 노드에 추가
	node := &Node{val: val}
	tail.next = node
	
	return node
}

//ㅁ-ㅁ-ㅁ : 가운데 노드의 레퍼런스 카운트를 0으로 만들어서 제거.
//(지우고자 하는 노드를 받고, 지우고자하는 노드가 나올 떄 까지 전진) : O(N)
func RemoveNode(node *Node, root *Node, tail *Node) (*Node ,*Node)  {

	if node == root {
		root = root.next
		if root == nil{
			tail = nil
		}
		return root,tail
	}

	prev := root
	//이전 노드의 next가 현재 노드가 아니면 찾을 떄까지 전진
	for prev.next != node{
		prev = prev.next
	}

	
	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next
	}
	
	return root,tail
}

func PrintNodes(root *Node){
	node := root
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}