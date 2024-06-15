package main
func main(){
	x :=[]int{1,2,2,3,4,5}
	for i,v := range x{
		println("位置", i,"值：", v)
	}
	
}
