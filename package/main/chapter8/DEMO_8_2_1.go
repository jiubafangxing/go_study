	
package chapter8
import(
	"log"
)
func DEMO_8_2_1test(){
	exit :=make(chan struct{})
	data := make(chan string)

	go func(){
		defer close(exit)
		d1 := <- data
		log.Println(d1+"from child ")

	}()

	data <- "hello world"
	<- exit
}

func DEMO_8_2_1main(){
	DEMO_8_2_1test()
}
