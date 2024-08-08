package chapter11

import (
	"github.com/jiubafangxing/go_study/package/main/chapter10"
	"log"
	"reflect"
)

func Chapter1118Test() {
	c := chapter10.C1011data{}
	ct := reflect.TypeOf(c)
	ctn := ct.NumField()
	for i := 0; i < ctn; i++ {
		ctf := ct.Field(i)
		log.Println(ctf)
	}
}
