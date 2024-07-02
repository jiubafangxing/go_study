package main

func main(){
	a := 1
	p := (*int)(&a)
	println(*p)
}
