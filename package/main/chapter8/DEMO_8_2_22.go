	
package chapter8
import(
	"time"
	"runtime"
)
func DEMO_8_2_22test(){
	ch := make(chan int)
	for i:=0;i<10;i++{
		go func(){
			<- ch
		}()
	}	
}

func DEMO_8_2_22main(){
	DEMO_8_2_22test()
	for{
		time.Sleep(time.Second)
		runtime.GC()
	}
}
