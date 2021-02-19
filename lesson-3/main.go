package main

import (
	"GB/lesson-3/prototype"
	"GB/lesson-3/singleton"
	"fmt"
)

func main() {
	s1 := singleton.NewSingleStruct()
	s2 := singleton.NewSingleStruct()

	fmt.Println(s1 == s2)

	p1 := prototype.NewPrototypeStruct()

	p2 := p1.Clone()

	fmt.Println(p1.Compare(p2))
}
