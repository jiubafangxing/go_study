	
package main
import(
	"log"
)
func test(){
	data := make(chan int)
	var send <- chan int = data
	var rece chan <- int = data
	//can not use rece to send msg
	//rece <- 1
	log.Println(rece)
	//can not use send to receive msg
	//errReceive := <- send
	//log.Println(errReceive)
	log.Println(send)

}

func main(){
	test()
}
