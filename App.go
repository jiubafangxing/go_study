package main

import (
	"log"
	"unsafe"

	"github.com/jiubafangxing/go_study/package/main/chapter10"
)

func main() {
	c := chapter10.BuildC1011Data()
	p := (*struct {
		x int
		y string
	})(unsafe.Pointer(c))
	p.y = "hello"
	log.Println(p.x)
	log.Println(p.y)
}
