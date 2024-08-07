package chapter1
func for_range_demomain(){
	x :=[]int{1,2,2,3,4,5}
	for i,v := range x{
		println("位置", i,"值：", v)
	}
	
}
