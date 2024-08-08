package chapter11
import (
"reflect"
"log"
)

func Chapter1116Test(){
	m := manager{
		level:1,
	}
	m.user.name = "john"
	m.user.age = 12
	mt := reflect.TypeOf(m)
	f,_ := mt.FieldByName("name")
	log.Printf("name : %s, type : %s\n", f.Name, f.Type)
	f = mt.FieldByIndex([]int{1,1})
	log.Printf("name : %s, type : %s\n", f.Name, f.Type)
}
