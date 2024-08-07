	
package main
import(
	"log"
	"sync"
)
type receiver struct{
	sync.WaitGroup
	data chan int
}
func test(){
	rece := &receiver{
		data: make(chan int),
	}
	rece.Add(1)
	go func(){
		defer rece.Done()
		for x := range rece.data{
			log.Println(x)
		}
	}()
	rece.data <- 1
	rece.data <- 2
	close(rece.data)
	rece.Wait()
}

func main(){
	test()
}
