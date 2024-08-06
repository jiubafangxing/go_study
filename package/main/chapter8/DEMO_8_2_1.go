	
package main
import(
	"log"
)
func test(){
	exit :=make(chan struct{})
	data := make(chan string)

	go func(){
		defer close(exit)
		d1 := <- data
		log.Println(d1+"from child ")

	}()

	data <- "hello world"
	<- exit
}

func main(){
	test()
}
