package main

import "fmt"

type T struct {
	hello int
}

func (t T) Hello() int {
	return t.hello
}

func (t *T) HelloGo() int {
	return t.hello
}

func (t *T) SetHello(a int) {
	t.hello = a
}

type W struct {
	T
	world int
}

func main() {

	t := T{hello: 7}
	t.Hello()
	t.HelloGo()
	t.SetHello(9)
	tx := &t
	tx.Hello()
	tx.HelloGo()
	tx.SetHello(10)

	w := &W{}
	w.hello = 98
	w.world = 99

	w.SetHello(10)

	fmt.Printf("%d    %d\n", w.HelloGo(), w.Hello())

}
