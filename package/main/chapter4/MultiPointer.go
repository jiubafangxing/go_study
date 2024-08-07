
package chapter4
func MultiPointertest(p **int){
	x :=1
	*p = &x
}

func MultiPointertest2(p *int){
	x :=2
	*p = x
}
func MultiPointermain(){
	var p *int
	MultiPointertest(&p)
	println(*p)
	MultiPointertest2(p)
	println(*p)
}
