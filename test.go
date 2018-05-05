package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := TestType{Hoge: 5, Fuga: "hogefuga", Typo: "TestType"}
	fmt.Println(t)
	fmt.Println(t.Foo())
	fmt.Println(t.Type())

	var i IF
	i = t

	fmt.Println(i)
	fmt.Printf("%#v\n", reflect.TypeOf(t))
	//r := i.(i.Type())
	//fmt.Println(r.Foo())
}

type IF interface {
	Type() string
}

type TestType struct {
	Hoge int
	Fuga string
	Typo string
}

func (s TestType) Foo() string {
	return fmt.Sprintln("Hello"+string(s.Hoge)+",", s.Fuga)
}

func (s TestType) Type() string {
	return s.Typo
}
