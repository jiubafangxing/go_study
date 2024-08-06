	
package main
import(
	"log"
)
func test(){
	a, b:= make(chan int ), make(chan int , 3)
	b <- 1
	b <- 2
	log.Printf("a len : %d, cap ; %d", len(a), cap(a))
	log.Printf("b len : %d, cap ; %d", len(b), cap(b))
	log.Println(<-b)
	log.Printf("b len : %d, cap ; %d", len(b), cap(b))

}

func main(){
	test()
}
