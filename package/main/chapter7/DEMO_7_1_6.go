package chapter7

import (
	"fmt"
	"log"
)

type Ner interface {
	a()
	b(int)
	c(string) string
}
type DEMO_7_1_6N int

func (n DEMO_7_1_6N) a() {
	log.Println(n)
}

func (n *DEMO_7_1_6N) b(a DEMO_7_1_6N) {
	*n = (*n) + a
}

func (n *DEMO_7_1_6N) c(b string) string {
	return fmt.Sprintf("%s-%d", b, *n)
}

func DEMO_7_1_6test() {
	var n DEMO_7_1_6N = 1
	var a DEMO_7_1_6N = 1
	n.a()
	n.b(a)
	log.Println(n)
	log.Println(n.c("hello"))
}

func DEMO_7_1_6main() {
	DEMO_7_1_6test()
}
