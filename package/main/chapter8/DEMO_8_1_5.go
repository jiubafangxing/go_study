	
package main
import(
	"log"
	"sync"
)
func test(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		wg.Wait()
		log.Println("child2 listen child1 done")
	}()

	go func(){
		wg.Done()
		log.Println("child done")
	}()	
	wg.Wait()
	log.Println("main listen child done")
}

func main(){
	test()
}
