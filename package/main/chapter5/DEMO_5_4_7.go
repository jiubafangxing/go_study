	
package chapter5
import(
	"time"
)
func DEMO_5_4_7test(){
	dict := map[int]int{
		1:2,
		2:2,
	}
	go func(){
		for{
			dict[1] += 1
			time.Sleep(time.Microsecond)			
		}
	}()

	go func(){
		for{
			_ = dict[1]
			time.Sleep(time.Microsecond)			
		}
	}()
	select {}


}

func DEMO_5_4_7main(){
	DEMO_5_4_7test()
}
