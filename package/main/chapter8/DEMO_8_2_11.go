	
package main
import(
	"log"
)
func test(){
	data := make(chan int)
	var send chan <- int = data
	var rece <- chan int = data
	
	//can not convert back
	//var data1 chan int = send
	//var data2 chan int = rece 

	log.Println(send)
	log.Println(rece)
}

func main(){
	test()
}
