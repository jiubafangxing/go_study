
package main
func test(p **int){
	x :=1
	*p = &x
}

func test2(p *int){
	x :=2
	*p = x
}
func main(){
	var p *int
	test(&p)
	println(*p)
	test2(p)
	println(*p)
}
