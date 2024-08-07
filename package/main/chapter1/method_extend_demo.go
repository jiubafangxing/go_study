package chapter1

import "fmt"

type Y int

func (y *Y) inc() {
	*y++
}

type XE struct {
	Y
	name string
}

func method_extend_demomain() {
	var xe XE
	xe.name = "a"
	xe.inc()
	fmt.Println(xe.Y)

}
