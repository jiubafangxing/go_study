	
package chapter8
import(
	"log"
	"runtime"
	"time"
)
func DEMO_8_1_20test(){

	go func(){
		for i:=0;i< 10;i++{
			time.Sleep(time.Millisecond)
			log.Println(i)
		}

	}()
	runtime.Goexit()
	log.Println("end")
}

func DEMO_8_1_20main(){
	DEMO_8_1_20test()
}
