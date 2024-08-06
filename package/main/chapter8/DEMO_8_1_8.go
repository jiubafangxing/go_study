	
package main
import(
	"log"
	"runtime"
)
func test(){
	//gosched
	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})
	go func(){
		defer close(exit)
		go func(){
			log.Println("b")
		}()

		for i:=0;i<10;i++{
			log.Println("a")
			if i ==1{
				runtime.Gosched()
			}
		}
	}()
	<-exit

}

func main(){
	test()
}
