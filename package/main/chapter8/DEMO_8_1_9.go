	
package chapter8
import(
	"log"
	"runtime"
)
func DEMO_8_1_9test(){
	exit := make(chan struct{})
	go func(){
		defer close(exit)
		defer log.Println(" a exit")
		func (){
			defer log.Println("b exit")
			func(){
				log.Println("c execute")
				runtime.Goexit()
				log.Println("c end")
			}()
		}()
	}()

	<- exit
	log.Println("DEMO_8_1_9main exit")

}

func DEMO_8_1_9main(){
	DEMO_8_1_9test()
}
