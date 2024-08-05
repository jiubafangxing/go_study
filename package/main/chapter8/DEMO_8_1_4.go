	
package main
import(
	"log"
	"sync"
)
func test(){
 	var wg  sync.WaitGroup
	go func(){
		wg.Add(1)
		log.Println("hi")
	}()
	wg.Wait()
	log.Println("exit")
}

func main(){
	test()
}
