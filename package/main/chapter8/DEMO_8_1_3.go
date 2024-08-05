	
package main
import(
	"log"
	"sync"
	"time"
)
func test(){
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
	println("main")
	wg.Wait()
	println("main exit")
}

func main(){
	test()
}
