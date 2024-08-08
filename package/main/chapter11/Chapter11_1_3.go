package chapter11

import (
	"fmt"
	"reflect"
)

type mybyte int8

func Chapter1113Test() {
	a := reflect.ArrayOf(3, reflect.TypeOf(mybyte(0)))
	b := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	fmt.Print(a, b)
}
