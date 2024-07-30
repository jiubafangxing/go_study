package main

func test(a ...int){
	for i := range a{
		println(i)
	}
}

func main(){
	c := [3]int{1,2,3}
	test(c[:]...)
}

