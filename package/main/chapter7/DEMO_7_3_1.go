package chapter7

import (
	"log"
)

type DEMO_7_3_1stringer interface {
	toString() string
}
type data struct {
	age int
}

func (d data) toString() string {
	return "a"
}

func DEMO_7_3_1test() {
	var d data = data{
		age: 1,
	}
	var d1 interface{} = d
	if xv, ok := d1.(DEMO_7_3_1stringer); ok {
		log.Println(xv)
	}
	if n, ok := d1.(data); ok {
		log.Println(n)
	}

}

func DEMO_7_3_1main() {
	DEMO_7_3_1test()
}
