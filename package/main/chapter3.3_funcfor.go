package main
func count() int{
	println("count")
	return 3
}
func main(){
	for x:= count();x< 10; x++{
		println(x);
	}
	for x:=1;x<count(); x++{
		println(x);
	}
}
