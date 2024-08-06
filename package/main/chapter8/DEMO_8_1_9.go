	
package main
import(
	"log"
	"runtime"
)
func test(){
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
	log.Println("main exit")

}

func main(){
	test()
}
