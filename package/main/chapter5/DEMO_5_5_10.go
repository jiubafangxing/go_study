	
package chapter5
import(
	"log"
)
func DEMO_5_5_10test(){
	a := make(chan struct{})

	go func(){
		log.Println("call ")
		a <- struct{}{}

	}()
	<- a
	log.Println("receive msg")
}

func DEMO_5_5_10main(){
	DEMO_5_5_10test()
}
