	
package main
import(
	"log"
)
func test(){
	var a,b chan int = make (chan int, 3),make (chan int)
	log.Println(a == b)
	var c chan int
	log.Println(c == b)
	
}

func main(){
	test()
}
