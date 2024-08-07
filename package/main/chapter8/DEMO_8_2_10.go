package chapter8
import(
	"log"
	"sync"
)
func DEMO_8_2_10test(){
	var wg  sync.WaitGroup
	wg.Add(2)
	ch1, ch2 := make(chan int), make(chan int)
	go func(){
		defer wg.Done()
		for{
			var (
				name string
				x int 
				ok bool
			)
			select{
			case x, ok = <-ch1 :
				name = "ch1"
			case x, ok = <-ch2 :
				name = "ch2"

			}
			if ! ok {
				return 
			}
			log.Println(name, x)
		}		
	}()
	go func(){
		defer wg.Done()
		defer close(ch1)
		defer close(ch2)
		for i:=0;i<10;i++{
			select{
			case ch1 <- i:
			case ch2 <- i*10:
			}
		}
	}()

	wg.Wait()
}

func DEMO_8_2_10main(){
	DEMO_8_2_10test()
}
