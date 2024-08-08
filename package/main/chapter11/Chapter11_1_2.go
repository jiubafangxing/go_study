package chapter11

import (
	"log"
	"reflect"
)

type Chapter1112DataA int
type Chapter1112DataB int

func Chapter1112Test() {
	a, c := Chapter1112DataA(1), Chapter1112DataA(2)
	b := Chapter1112DataB(3)
	ta, tc, tb := reflect.TypeOf(a), reflect.TypeOf(c), reflect.TypeOf(b)
	log.Println(ta.Name() == tc.Name())
	log.Println(ta.Kind() == tb.Kind())
}
