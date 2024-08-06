	
package main
import(
	"log"
	"sync"
)
func test(){
	var wg sync.WaitGroup
	start := make(chan struct{})
	for i:=0;i<8;i++{
		wg.Add(1)
		go func(a int){
			defer wg.Done()
			log.Printf("id:%d ready!\n", a)
			<-start
			log.Printf("id:%d go\n",a)
		}(i)
	}

	log.Println("ready go")
	close(start)
	wg.Wait()
}

func main(){
	test()
}
