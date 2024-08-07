	
package chapter8
import(
	"log"
	"time"
)


func DEMO_8_1_2test(){
	exit := make(chan struct{})
	go func(){
		log.Println("child")
		time.Sleep(time.Second)
		close(exit)
	}()
	log.Println("DEMO_8_1_2main")
	<-exit
	log.Println("end")
		
}

func DEMO_8_1_2main(){
	DEMO_8_1_2test()
}
