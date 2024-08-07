package chapter1

import "fmt"

type parent struct {
	parent_name string
}

type child struct {
	parent
	child_name string
}

func struct_demomain() {
	var a child
	a.parent_name = "li"
	a.child_name = "wang"
	println(a.child_name, a.parent_name)
	fmt.Println(a)
}
