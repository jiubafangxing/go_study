package chapter9

import (
	"cmp"
	"log"
)

func maxValue[T cmp.Ordered](x, y T) T {
	if x > y {
		return x
	} else {
		return y
	}
}
func Chapter9_1_1Test() {
	x, y := 1, 2
	z := maxValue(x, y)
	log.Printf("max value %d\n", z)

	x1, y1 := 'a', 'b'
	z1 := maxValue(x1, y1)
	log.Printf("max value %c", z1)
	return
}
