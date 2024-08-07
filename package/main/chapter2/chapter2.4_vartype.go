package chapter2

import (
	"fmt"
)

func chapter24_vartypemain() {
	var a, b, c = 100, 0144, 0x64
	fmt.Println(a, b, c)
	fmt.Printf("num a %#o", a)
	fmt.Printf("num a %o", a)
	fmt.Printf("num b %#x", b)
	fmt.Printf("num b %x", b)
	fmt.Printf("num c %b", c)
}
