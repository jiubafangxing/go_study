package chapter3
func chapter3_3_copy_slicedemomain(){
	var arr [3]int  = [3]int{1,2,3}
	for i,v := range arr{
		println(i,v)
	}
	for i1,_ := range arr[0:2]{
		arr[i1] = 0
	}
	for i,v := range arr{
		println(i,v)
	}
}
