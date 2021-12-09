package main

import "fmt"

func main() {
	var i int
	for {
		if i == 5 {
			i++
			continue //건너 뜀
		}
		if i == 6 {
			break // 빠져나감
		}
		fmt.Println(i)
		i++
	}
}