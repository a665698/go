package main

import (
	"fmt"
)

type THandle interface {
	Add()
}

type Menu struct {
	Type string
	Name string
	Key string
	THandle
}

type Action1 struct {
	Menu
	aa string
}
type Action2 struct {
	Menu
	aa string
}

func (a *Action1) Add()  {
	a.aa = "1233"
	fmt.Printf(a.aa)
}

func (m *Menu) Add()  {
	fmt.Println("11111")
	fmt.Printf("%T", m)
	fmt.Println(m)
}

func main() {
	a := &Menu{THandle: &Action2{}}
	a.THandle.Add()
	fmt.Println(a)
	fmt.Printf("%T %+v", a.THandle, a.THandle)
}



