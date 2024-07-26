package main
func main(){
	a := "thlias"
	for i, c :=range a {
		println(i, c)
	}
	println("...")
	for i := range a{
		println(i, a[i])
	}
	println("...")
	for _, c := range a{
		println( c)
	}
	println("...")
	for i,_ := range a{
		println(i, a[i])
	}
}
