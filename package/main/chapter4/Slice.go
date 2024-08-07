package chapter4

func Slicetest(a ...int){
	for i := range a{
		println(i)
	}
}

func Slicemain(){
	c := [3]int{1,2,3}
	Slicetest(c[:]...)
}

