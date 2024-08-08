package chapter11
import (
	"log"
	"reflect"
)
type C1132Data int 
func (C1132Data) Add(a ...interface{})(int){
	result := 0
	for _,num := range a{
		nv := reflect.ValueOf(num)
		if nv.Kind() == reflect.Int{
			result += int(nv.Int())
		}else{
			return -1
		}	
	}
	return result
}
func Chapter1132Test(){
	c := C1132Data(1)
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

	result2 := m.CallSlice([]reflect.Value{
		reflect.ValueOf([]interface{}{1,3,4,5,6}),
	})
	for _,v2 := range result2{
		log.Println(v2)
	}
}
