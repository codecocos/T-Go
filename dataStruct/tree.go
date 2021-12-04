package dataStruct

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

func (t *Tree) DFS1() {
	DFS1(t.Root)
}

//재귀호출 방식
func DFS1(node *TreeNode) {
	fmt.Printf("%d->",node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}
}

//스택 방식
func (t *Tree) DFS2() {
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s)>0{
		var last *TreeNode
		//맨 마지막에 있는 값을 가져오고 , 맨 마지막을 없앤 슬라이스를 대입 
		last, s = s[len(s)-1],s[:len(s)-1]
		
		fmt.Printf("%d->", last.Val)

		// for i := 0; i < len(last.Childs); i++ { // 1->4->10->9->3->8->7->2->6->5->
		// 	s = append(s, last.Childs[i])
		// }

		for i :=  len(last.Childs)-1; i >= 0 ; i-- { // 1->2->5->6->3->7->8->4->9->10->
			s = append(s, last.Childs[i])
		}
	}
}

func (t *Tree) BFS() {
	//큐를 만들고
	queue := []*TreeNode{}
	//루트를 큐에 넣고
	queue = append(queue, t.Root)

	//큐가 빌 때 까지 돈다.
	for len(queue) > 0 {
		var first *TreeNode
		//first 에는 큐의 첫번째, 큐에는 기존 큐의 두번째 부터
		first, queue = queue[0], queue[1:]

		fmt.Printf("%d->", first.Val)

	//꺼내 첫번째에 모든 자식들을 다시 큐에 집어 넣음 
	for i := 0; i < len(first.Childs); i++ {
		queue = append(queue, first.Childs[i])
	}
	}
}