	
package chapter8
import(
	"log"
	"sync"
)
func DEMO_8_2_9test(){
	c := make(chan int)
	var send chan <- int = c
	var rece <- chan int = c
	var wg sync.WaitGroup	
	wg.Add(2)
	go func(){
		defer wg.Done()
		defer close(c)
		for i:=0;i<3;i++{
			send <- i
		}
	}()

	go func(){
		defer wg.Done()
		for i := range  rece{
			log.Printf("child receive %d\n",i )
		}
	}()
	wg.Wait()
}

func DEMO_8_2_9main(){
	DEMO_8_2_9test()
}
