package chapter4

func EscapeHeaptest(p *int){
	go func(){
		println(*p)
	}()
}

func EscapeHeapmain(){
	a :=1
	EscapeHeaptest(&a)
}
