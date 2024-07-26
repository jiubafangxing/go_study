package main
type myt byte
const (
	a myt = 1 << iota
	b
	c
)
func main(){
	println(a)
	println(b)
	println(c)
}

