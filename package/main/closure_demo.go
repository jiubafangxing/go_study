package main

func genClosureFunc(x string) func (){
	return  func (){
		println(x)
	}
}

func main(){
	fa :=genClosureFunc("a")
	fb :=genClosureFunc("b")
	fa()
	fb()
	fa()
}
