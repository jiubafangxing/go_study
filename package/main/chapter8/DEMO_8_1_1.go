	
package chapter8
import(
	"log"
	"time"
)
func DEMO_8_1_1test(){
	go log.Println("hello")
	log.Println("hello DEMO_8_1_1main")
	time.Sleep(time.Second * 3)
}

func DEMO_8_1_1main(){
	DEMO_8_1_1test()
}
