	
package chapter8
import(
	"log"
	"sync"
	"time"
)
func DEMO_8_1_3test(){
	var wg  sync.WaitGroup
	for i:=0;i<10; i++{
		wg.Add(1)
		x := i
		go func(){
			defer wg.Done()
			time.Sleep(time.Second)
			log.Println(x)	
		}()
	}	
	println("DEMO_8_1_3main")
	wg.Wait()
	println("DEMO_8_1_3main exit")
}

func DEMO_8_1_3main(){
	DEMO_8_1_3test()
}
