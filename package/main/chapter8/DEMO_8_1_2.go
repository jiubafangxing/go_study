	
package main
import(
	"log"
	"time"
)


func test(){
	exit := make(chan struct{})
	go func(){
		log.Println("child")
		time.Sleep(time.Second)
		close(exit)
	}()
	log.Println("main")
	<-exit
	log.Println("end")
		
}

func main(){
	test()
}
