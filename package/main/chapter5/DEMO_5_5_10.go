	
package main
import(
	"log"
)
func test(){
	a := make(chan struct{})

	go func(){
		log.Println("call ")
		a <- struct{}{}

	}()
	<- a
	log.Println("receive msg")
}

func main(){
	test()
}
