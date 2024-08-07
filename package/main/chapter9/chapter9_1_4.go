package chapter9

import (
	"log"
)

type C914_A struct {
	x int
}
type C914_B struct {
	x string
}

func (c C914_B) c914print() {
	log.Println( c.x)
}
func (c C914_A) c914print() {
	log.Println( c.x)
}

type C914_AB interface {
	C914_A | C914_B
	c914print()
}

func print914[T C914_AB](t T) {
	t.c914print()
}
func Chapter914Test() {
	c1 := C914_A{
		x: 1,
	}
	c2 := C914_B{
		x: "hello ",
	}
	print914(c1)
	print914(c2)
	//string does not satisfy C914_AB (missing method c914print)
	//	print914("a")
	return
}
