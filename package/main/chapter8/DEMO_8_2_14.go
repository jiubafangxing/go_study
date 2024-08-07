	
package chapter8
import(
	"log"
	"time"
)
func DEMO_8_2_14test(){
	done := make(chan struct{})
	data := make(chan  int)
	go func(){
		defer close(done)
		for{
			select{
				case d1,ok := <- data:
					if !ok {
						return
					}
					log.Printf("data receive :%d", d1)
				default:

			}
			log.Println(time.Now())
			time.Sleep(time.Second)

		}
	}()


	time.Sleep(time.Second * 5)	
	data <- 1
	close(data)
	<- done
}

func DEMO_8_2_14main(){
	DEMO_8_2_14test()
}
