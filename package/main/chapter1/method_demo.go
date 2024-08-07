package chapter1

type X int

func (x *X) inc() {
	*x++
}

func method_demomain() {
	var x X
	x = 0
	x.inc()
	println(x)
}
