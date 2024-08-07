package chapter7

import (
	"log"
)

type DEMO_7_1_1tester interface {
	DEMO_7_1_1test()
	genStr() string
}
type chpter7N int

func (n chpter7N) DEMO_7_1_1test() {
	log.Println(n)
}
func (n *chpter7N) genStr() string {
	return "adf"
}

func DEMO_7_1_1main() {
	var n chpter7N = 1
	n.DEMO_7_1_1test()
	s := n.genStr()
	log.Println(s)
	var t DEMO_7_1_1tester = &n
	t.DEMO_7_1_1test()
}
