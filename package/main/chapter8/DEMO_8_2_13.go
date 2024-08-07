	
package chapter8
import(
	"log"
	"sync"
)
func DEMO_8_2_13test(){
	var wg sync.WaitGroup
	wg.Add(2)
	ch1 := make(chan int)
	go func(){
		defer wg.Done()
		for{
			select{
				case x ,ok := <- ch1:
					if ! ok{
						ch1 = nil
						break
					}
					log.Printf("分之1-%d\n",x)
				case x ,ok := <- ch1:
					if ! ok{
						ch1 = nil
						break
					}
					log.Printf("分之2-%d\n",x)
			
			}
			if nil == ch1{
				return
			}
		}
	}()

	go func(){
		defer wg.Done()
		defer close(ch1)
		for i:=0;i< 10;i++{
			select{

				case ch1 <- i:
				case ch1 <- i*100:

			}
		}
	}()
	wg.Wait()
}

func DEMO_8_2_13main(){
	DEMO_8_2_13test()
}
