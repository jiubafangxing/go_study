package chapter11
import (
"reflect"
"log"
"fmt"
)

type C11110DATA int



func Chapter11110Test(){
	x := C11110DATA(1)
	xt := reflect.TypeOf(x)
	st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	log.Println(xt.Implements(st))	
	numT := reflect.TypeOf(0)
	log.Println(xt.ConvertibleTo(numT))
	log.Println(xt.AssignableTo(numT), xt.AssignableTo(st))	
}
