package chapter2

import (
	"fmt"
)

func chapter24_floatposmain() {
	var a float32 = 1.1234567899
	var b float32 = 1.123345681
	var c float32 = 1.12334568
	fmt.Println(a == b)
	fmt.Println(a == c)
}
