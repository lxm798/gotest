package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Value int
}

type B interface {
	add(b *A)
	print()
}

type C struct {
	b B
	c *A
}

/*
func (a A) add(b A) {
	fmt.Printf(" non pointer at all\n")
	a.Value += b.Value
	b.Value += 5
}
*/

func (a *A) add(b *A) {
	fmt.Printf("canshu pointer\n")
	a.Value += b.Value
	b.Value += 1
}
func (a *A) inc(b *A) {
	fmt.Printf("canshu pointer \n")
	(*a).Value += (*b).Value
}
func (a *A) print() {
	fmt.Printf("aa %d\n", a.Value)
}

func (c *C) add() {
	fmt.Printf("type C:%s\n", reflect.TypeOf(c))
	fmt.Printf("type *C:%s\n", reflect.TypeOf(*c))
	fmt.Printf("type *C.C:%s\n", reflect.TypeOf(*c.c))
	fmt.Printf("type (*C).c:%s\n", reflect.TypeOf((*c).c))
	fmt.Printf("type *(C.c):%s\n", reflect.TypeOf(*(c.c)))
	var a A
	a.Value = 2
	fmt.Printf("type a:%s\n", reflect.TypeOf(a))
	fmt.Printf("type a:%s\n", reflect.TypeOf(&a))
	fmt.Printf("address of a %p\n", &a)
	c.c = &a
	fmt.Printf("address of c.c %p\n", c.c)
	fmt.Printf("type c.c:%s\n", reflect.TypeOf(c.c))
}
func main() {

	c := new(C)
	a := &A{Value: 3}
	c.c = a
	var e A
	f := A{Value: 6}
	c.b = a
	b := &A{Value: 4}
	c.b.add(b)
	c.b.print()
	a.print()
	d := c.b
	d.add(b)
	c.b.print()
	d.print()
	fmt.Println("type a", reflect.TypeOf(a))
	fmt.Println("type b", reflect.TypeOf(b))

	fmt.Println("type a", reflect.TypeOf(c))
	fmt.Println("type b", reflect.TypeOf(d))
	fmt.Println("type a", reflect.TypeOf(e))
	fmt.Println("type b", reflect.TypeOf(f))
	c.add()
	fmt.Printf("c.c %d\n", c.c.Value)
	var g A
	g.Value = 4
	fmt.Printf("address of c.c %p\n", c.c)
	fmt.Printf("address of c.c %p\n", &g)
	fmt.Printf("type a:%s\n", reflect.TypeOf(g))
	/*
		a := &A{Value: 3}
		b := A{Value: 4}
		a.add(b)
		fmt.Println("type a", reflect.TypeOf(a))
		fmt.Println("type b", reflect.TypeOf(b))
		fmt.Printf("result is %d,%d\n", a.Value, b.Value)
	*/
	is := [4]int{}
	fmt.Printf("\n1:%d\n", is[0])
	is1 := is
	is[0] = 4
	fmt.Printf("agein 1:%d\n", is[0])
	fmt.Printf("copy 1:%d\n", is1[0])

	eis := [4]A{}
	fmt.Printf("\neis 1:%d\n", eis[0].Value)
	fis := eis
	eis[0].Value = 200
	fmt.Printf("eis agein 1:%d\n", eis[0].Value)
	fmt.Printf("fis copy 1:%d\n", fis[0].Value)

	ais := make([]A, 4)
	fmt.Printf("init with not nil:%d\n", ais[0].Value)
	bis := ais
	ais[0].Value = 100
	fmt.Printf("ais index0:%d\n", ais[0].Value)
	fmt.Printf("bis index0:%d\n", bis[0].Value)
	ais[1].Value = 200
	fmt.Printf("ais index1:%d\n", ais[1].Value)
	fmt.Printf("bis index1:%d\n", bis[1].Value)
	aa := A{0}
	gis := append(ais, aa)
	aa.Value = 10
	ais[0].Value = 14532
	gis[0].Value = 10003223
	bis[0].Value = 2323
	fmt.Printf("----------painc---\n")
	fmt.Printf("ais indexlast:%d\n", ais[len(ais)-1].Value)
	fmt.Printf("bis indexlast:%d\n", bis[len(bis)-1].Value)
	fmt.Printf("ais index0:%d\n", ais[0].Value)
	fmt.Printf("bis index0:%d\n", bis[0].Value)
	fmt.Printf("gis index0:%d\n", gis[0].Value)
	ais[len(ais)-1].Value = 210
	bis = append(bis, A{200000})
	fmt.Printf("ais indexlast:%d,len:%d\n", ais[len(ais)-1].Value, len(ais))
	fmt.Printf("bis indexlast:%d, len:%d\n", bis[len(bis)-1].Value, len(bis))
	ais[0].Value = 1000
	fmt.Printf("ais index0 change 1000:%d\n", ais[0].Value)
	fmt.Printf("bis index0 change 1000:%d\n", bis[0].Value)

	cis := make([]int, 4)
	fmt.Printf("init with cis not nil:%d\n", cis[0])
	dis := ais
	cis[0] = 100
	fmt.Printf("cis index0:%d\n", cis[0])
	fmt.Printf("dis index0:%d\n", dis[0])
	cint := 3
	cis = append(cis, cint)
	cint = 4
	fmt.Printf("cis indexlast:%d\n", cis[len(ais)-1])

	var ch chan int
	if ch == nil {
		fmt.Printf("chan is nil\n")
	} else {
		fmt.Printf("chan is not nil\n")
	}
}
