	
package chapter8
import(
	"log"
	"sync"
)
func DEMO_8_1_5test(){
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
	log.Println("DEMO_8_1_5main listen child done")
}

func DEMO_8_1_5main(){
	DEMO_8_1_5test()
}
