package chapter11

import (
	"log"
	"reflect"
)

func Chapter1121Test() {
	a :=  1
	va, vp := reflect.ValueOf(a),reflect.ValueOf(&a).Elem()
	log.Println(va.CanAddr(),va.CanSet())
	log.Println(vp.CanAddr(),vp.CanSet())
}
