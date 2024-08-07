	
package chapter8
import(
	"log"
)
func DEMO_8_2_8test(){
	//buffer data from chan
	data := make(chan int,2)
	data <- 1
	data <- 2
	close(data)
	//if close chan more than one times , panic
	//close(data)
	for{
	   v,ok:=<-data
	   if !ok{
		log.Println("no data ")   
		break
	   }else{
		log.Println(v)
	   }

	}
}

func DEMO_8_2_8main(){
	DEMO_8_2_8test()
}
