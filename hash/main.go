package main

import (
	"fmt"

	"github.com/T-Go/dataStruct"
)

func main() {
	fmt.Println("abcde =", dataStruct.Hash("abcde"))
	fmt.Println("abcde =", dataStruct.Hash("abcde"))
	fmt.Println("tbcde =", dataStruct.Hash("tbcde"))
	fmt.Println("abcdf =", dataStruct.Hash("abcdf"))
	fmt.Println("abcdffdfsdfds =", dataStruct.Hash("abcdffdfsdfds"))

	m := dataStruct.CreateMap()
	m.Add("AAA","0107777777")
	m.Add("BBB","0108888888")
	m.Add("CDEFGTEFVDF","0111111111")
	m.Add("CCC","0177575757")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("BBB = ", m.Get("BBB"))
	fmt.Println("CDEFGTEFVDF = ", m.Get("CDEFGTEFVDF"))
	fmt.Println("CCC = ", m.Get("CCC"))
	fmt.Println("DDD = ", m.Get("DDD"))
}