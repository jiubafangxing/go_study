	
package chapter8
import(
	"log"
)
func DEMO_8_2_6test(){
	exit := make(chan struct{})
	data := make(chan int)
	go func(){
		defer close(exit)
		for i := range data{
			log.Printf("child receive msg %d\n", i)
		}
	}()
	data <- 1
	data <- 2
	data <- 3
	// if not close deadlock
	//close(data)
	<-exit
}

func DEMO_8_2_6main(){
	DEMO_8_2_6test()
}
