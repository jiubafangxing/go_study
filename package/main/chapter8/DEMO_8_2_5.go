
package main
import(
	"log"
)
func test(){
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

func main(){
	test()
}
