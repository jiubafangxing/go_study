package chapter7

import (
	"log"
)

type DEMO_7_1_4tester interface {
	mystringer
	DEMO_7_1_4test()
}
type mystringer interface {
	toString() string
}
type DEMO_7_1_4N int

func (n *DEMO_7_1_4N) DEMO_7_1_4test() {
	log.Println(*n)
}

func (n *DEMO_7_1_4N) toString() string {
	return "a"
}

func DEMO_7_1_4test1() {
	var n DEMO_7_1_4N = 1
	var DEMO_7_1_4tester1 DEMO_7_1_4tester = &n
	DEMO_7_1_4tester1.DEMO_7_1_4test()
	log.Println(DEMO_7_1_4tester1.toString())
	var str1 mystringer = DEMO_7_1_4tester1
	log.Println(str1.toString())
	//can not be used on  this way
	//var DEMO_7_1_4tester2 DEMO_7_1_4tester = str1
	//DEMO_7_1_4tester2.DEMO_7_1_4test()
}

func DEMO_7_1_4main() {
	DEMO_7_1_4test1()
}
