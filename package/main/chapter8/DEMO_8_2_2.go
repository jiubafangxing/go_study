	
package main
import(
	"log"
)
func test(){
	c := make(chan string,2)
	c <- "1"
	c <- "2"
	log.Println(<-c)
	log.Println(<-c)
}

func main(){
	test()
}
