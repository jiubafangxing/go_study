	
package main
import(
	"log"
	"runtime"
	"time"
)
func test(){

	go func(){
		for i:=0;i< 10;i++{
			time.Sleep(time.Millisecond)
			log.Println(i)
		}

	}()
	runtime.Goexit()
	log.Println("end")
}

func main(){
	test()
}
