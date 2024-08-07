
package chapter8
import(
	"log"
	"time"
)
func DEMO_8_2_15test(){
	done := make(chan struct{})
	data := []chan int{
		make(chan int,3),
	}
	go func(){
		defer close(done)
		time.Sleep(time.Second)
		for i:=0;i<10;i++{
			select {

			case data[len(data)-1] <- i:
			default:
				data = append(data, make(chan int ,3))

			}

		}
	}()

	<- done
//	for i:=0;i< len(data);i++{
//		dataitem := data[i]
	for _,dataitem := range data{
		close(dataitem)	
		for dataitemd := range dataitem{
			log.Printf("receive data %d\n",dataitemd)
		}
	}
}

func DEMO_8_2_15main(){
	DEMO_8_2_15test()
}
