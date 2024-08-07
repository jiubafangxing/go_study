package chapter1

func genClosureFunc(x string) func (){
	return  func (){
		println(x)
	}
}

func closure_demomain(){
	fa :=genClosureFunc("a")
	fb :=genClosureFunc("b")
	fa()
	fb()
	fa()
}
