package main
	func errorConvert(){
		x :=100
		p :=*(int(&x))
		println(p)
	}


	func main(){
		errorConvert();
	}

