package chapter3
import (
"fmt"
)
func chapter3_3outtermain(){
	var a string = "hello"
	for _,v := range a{
		fmt.Println("")
		fmt.Printf("%c",v)
	}
}
