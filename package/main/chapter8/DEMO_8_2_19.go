	
package main
import(
	"log"
	"time"
	"os"
)
func test(){
	go func(){
		for{
			select{
				case <- time.After(5* time.Second):
					log.Println("time to end")
					os.Exit(0)
			}
		}
	}()
	go func(){
		t := time.Tick(time.Second)
		for{
			select{
			case i:=<-t:
				log.Println(i)

			}

		}
	}()
	<- (chan struct{})(nil)
}

func main(){
	test()
}
