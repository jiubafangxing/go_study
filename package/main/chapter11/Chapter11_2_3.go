package chapter11
import (
	"log"
	"reflect"
)
type C1123User struct{
	Name string
	Age int
} 

func Chapter1123Test(){
	user :=C1123User{
		Name:"laoli",
		Age:21,
	}
	v :=reflect.ValueOf(&user)
	if !v.CanInterface(){
		log.Println("")
	}
	iv,ok := v.Interface().(*C1123User)
	if !ok{
		log.Fatalln("Interface err")
	}
	iv.Age++
	log.Println(v)
}
