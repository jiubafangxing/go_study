package chapter11
import (
	"log"
	"reflect"
)
type C1131Data int 
func (C1131Data) Add(a,b int)(int){
	return a+b
}
func Chapter1131Test(){
	c := C1131Data(1)
	cv := reflect.ValueOf(c)
	m := cv.MethodByName("Add")
	params :=[]reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}
	result := m.Call(params)
	for _,v := range result{
		log.Println(v)
	}
}
