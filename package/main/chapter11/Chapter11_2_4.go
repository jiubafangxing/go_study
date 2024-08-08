package chapter11
import (
	"log"
	"reflect"
)

func Chapter1124Test(){
	c :=make(chan int, 4)
	v :=reflect.ValueOf(c)
	if v.TrySend(reflect.ValueOf(100)){
		log.Println(v.TryRecv())
	}
}
