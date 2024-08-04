	
package main
import(
	"time"
)
func test(){
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

func main(){
	test()
}
