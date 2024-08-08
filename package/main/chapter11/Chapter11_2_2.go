package chapter11

import (
	"log"
	"reflect"
	"unsafe"
)

type C1122User struct {
	Name string
	Age  int
}

func Chapter1122Test() {
	c := C1122User{
		Name: "x1",
		Age:  2,
	}
	elem := reflect.ValueOf(&c).Elem()
	nf := elem.FieldByName("Name")
	if nf.CanSet() {
		nf.SetString("y1")
	}
	af := elem.FieldByName("Age")
	if af.CanAddr() {
		*(*int)(unsafe.Pointer(af.UnsafeAddr())) = 12
	}
	log.Println(c)
}
