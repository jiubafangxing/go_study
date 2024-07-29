package main

func test(p *int){
	go func(){
		println(*p)
	}()
}

func main(){
	a :=1
	test(&a)
}
