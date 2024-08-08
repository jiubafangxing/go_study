package chapter11

import (
	"log"
	"reflect"
)

type Chapter1111Data int

func Chapter1111Test() {
	a := Chapter1111Data(1)
	asource := reflect.TypeOf(a)
	log.Printf("real type is %s, parent type is %s", asource.Name(), asource.Kind())
}
