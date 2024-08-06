	
package main
import(
	"log"
)
func test(){
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

func main(){
	test()
}
