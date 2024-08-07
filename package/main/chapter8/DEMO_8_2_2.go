	
package chapter8
import(
	"log"
)
func DEMO_8_2_2test(){
	c := make(chan string,2)
	c <- "1"
	c <- "2"
	log.Println(<-c)
	log.Println(<-c)
}

func DEMO_8_2_2main(){
	DEMO_8_2_2test()
}
