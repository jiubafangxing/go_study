	
package main
import(
	"log"
	"fmt"
)
type stringer interface{
	toString()string
}
type data struct{
	age int
}

func (d data)toString()string{
	return "a" 
}


func test(){
	var d data = data{
		age:1
	}
	var d1 interface{} = d
	if xv, ok := d1.(stringer);ok{
		log.Println(xv)
	}
}

func main(){
	test()
}
