package chapter11
import (
	"log"
	"unsafe"
)

func Chapter1126Test(){
	var b interface{}= (*int)(nil)
	iface :=(*[2]uintptr)(unsafe.Pointer(&b))
	log.Println(iface,iface[1] == 0)
}
