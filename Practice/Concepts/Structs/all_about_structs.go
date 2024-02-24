package main

import "fmt"

type S struct {
	a int
}

func (s S) Foo() {
	s.a = 100
	fmt.Println(s.a)
}

func (s *S) Foos() {
	s.a = 200
	fmt.Println(s)
}

func main() {
	s := S{
		a: 10,
	}
	s.Foo()
	fmt.Println(s)
	s.Foos()
	fmt.Println(s)

	fmt.Println("with pointer")
	p := new(S)
	p.a = 30
	p.Foo()
	fmt.Println(p)
	p.Foos()
	fmt.Println(p)
}
