package chapter3
func chapter3_3_sliceandrangemain(){
	var arr [6]int =[6]int{1,2,3,4,5,6}
	for i,_  := range arr{
		arr[i] = 0
	}
	for i,_ := range arr{
		println(arr[i]) 
	}
	
}
