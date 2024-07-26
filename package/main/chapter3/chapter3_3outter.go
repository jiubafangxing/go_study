package main
import (
"fmt"
)
func main(){
	var a string = "hello"
	for _,v := range a{
		fmt.Println("")
		fmt.Printf("%c",v)
	}
}
