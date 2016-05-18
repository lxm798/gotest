package main

import (
	"fmt"
)

type A struct {
	Value int
}

func (a A) add(b A) {
	fmt.Printf(" non pointer at all\n")
	a.Value += b.Value
}

func (a A) add(b *A) {
	a.Value.Printf("canshu pointer\n")
	a.Value += b.Value
}
func main() {
	a := &A{Value: 3}
	b := &A{Value: 4}
	a.add(b)
	fmt.Printf("result is %d\n", a.Value)
}
