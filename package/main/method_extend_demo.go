package main

import "fmt"

type Y int

func (y *Y) inc() {
	*y++
}

type XE struct {
	Y
	name string
}

func main() {
	var xe XE
	xe.name = "a"
	xe.inc()
	fmt.Println(xe.Y)

}
