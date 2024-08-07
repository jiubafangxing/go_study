package chapter9

import (
	"fmt"
)

func chapter915print[T int | string](t T) {
	fmt.Println(t)
}

func Chapter915Test() {
	chapter915print(1)
	chapter915print("abc")
//	chapter915print(true)
	return
}
