package main

type X int

func (x *X) inc() {
	*x++
}

func main() {
	var x X
	x = 0
	x.inc()
	println(x)
}
