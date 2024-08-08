package chapter11

import (
	"log"
	"reflect"
)

func Chapter1114Test() {
	a := []int{1, 2, 3}
	atype := reflect.TypeOf(a)
	log.Println(atype.Elem())
}
