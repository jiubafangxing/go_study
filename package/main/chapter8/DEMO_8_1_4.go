	
package chapter8
import(
	"log"
	"sync"
)
func DEMO_8_1_4test(){
 	var wg  sync.WaitGroup
	go func(){
		wg.Add(1)
		log.Println("hi")
	}()
	wg.Wait()
	log.Println("exit")
}

func DEMO_8_1_4main(){
	DEMO_8_1_4test()
}
