package chapter11
import (
	"log"
	"reflect"
)

func Chapter1125Test(){
	var a interface{} = nil
	var b interface{} = (*int)(nil)
	log.Println(a == nil)
	log.Println(b == nil, reflect.ValueOf(b).IsNil())
}
