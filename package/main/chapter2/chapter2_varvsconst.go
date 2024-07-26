package main
var x = 0x100
const y = 0x200
func main(){
	println(&x, x)
	println(&y, y)
}
