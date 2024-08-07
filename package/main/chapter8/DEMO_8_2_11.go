	
package chapter8
import(
	"log"
)
func DEMO_8_2_11test(){
	data := make(chan int)
	var send chan <- int = data
	var rece <- chan int = data
	
	//can not convert back
	//var data1 chan int = send
	//var data2 chan int = rece 

	log.Println(send)
	log.Println(rece)
}

func DEMO_8_2_11main(){
	DEMO_8_2_11test()
}
