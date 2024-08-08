package chapter11
import (
	"log"
	"reflect"
)

func Chapter1127Test(){
	v := reflect.ValueOf(struct{name string}{})
	log.Println(v.FieldByName("name").IsValid())
	log.Println(v.FieldByName("name1").IsValid())
}
