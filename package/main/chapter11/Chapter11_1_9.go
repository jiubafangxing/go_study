package chapter11
import (
"reflect"
"log"
"github.com/jiubafangxing/go_study/package/main/chapter10"
)


func Chapter1119Test(){
	c:=	chapter10.C1011data{}
	ct := reflect.TypeOf(c)
	ctn := ct.NumField()
	for i:=0;i< ctn;i++{
		ctf :=ct.Field(i)
		name := ctf.Tag.Get("field")
		log.Println(ctf)
		log.Println(name)
	}
}
