package main

import "fmt"

type parent struct {
	parent_name string
}

type child struct {
	parent
	child_name string
}

func main() {
	var a child
	a.parent_name = "li"
	a.child_name = "wang"
	println(a.child_name, a.parent_name)
	fmt.Println(a)
}
