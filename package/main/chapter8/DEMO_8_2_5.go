
package chapter8
import(
	"log"
)
func DEMO_8_2_5test(){
	exit := make(chan struct{})
	data := make (chan int)
	go func(){
		defer close(exit)
		for{
			x,ok := <- data

			if !ok{
				return
			}
			log.Printf("child receive %d", x)
		}
	}()
	data <- 1
	data <- 2
	close(data)
	<- exit
}

func DEMO_8_2_5main(){
	DEMO_8_2_5test()
}
