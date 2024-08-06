
package main
import(
	"log"
	"sync"
)
func test(){
	var  wg sync.WaitGroup
	wg.Add(2)
	ch1 ,ch2 := make(chan int), make(chan int)
	go func(){
		defer wg.Done()
		for{
			select {
			case x ,ok := <- ch1:
				if !ok{
					ch1 = nil
					break
				}
				log.Println(x)
			case x, ok := <- ch2:
				if !ok{
					ch2= nil
					break
				}
				log.Println(x)
			}
			if nil == ch1 && nil == ch2{
				return
			}
		}

	}()
	go func(){
		defer wg.Done()
		defer close(ch1)
		defer close(ch2)
		for i:=0;i< 10;i++{
			select {
			case  ch1 <- i:
			case ch2 <- i*10:
			}
		}
	}()
	wg.Wait()
}

func main(){
	test()
}
