	
package main
import(
	"log"
	"time"
)
func test(){
	go log.Println("hello")
	log.Println("hello main")
	time.Sleep(time.Second * 3)
}

func main(){
	test()
}
