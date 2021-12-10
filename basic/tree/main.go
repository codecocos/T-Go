package main

import (
	"fmt"

	"github.com/T-Go/basic/dataStruct"
)

func main() {
	tree := dataStruct.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}

	tree.DFS1(); // 1->2->5->6->3->7->8->4->9->10->
	fmt.Println()
	tree.DFS2()
	fmt.Println()
	
	tree.BFS()

}

/*
 ____1____
|    |    |
2    3    4
|\   |\   |\
5 6  7 8  9 10
*/
