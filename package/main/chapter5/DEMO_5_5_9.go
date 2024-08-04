	
package main
import(
	"log"
)
func test(){
	var a = struct{}{}
	var b = [3]struct{}{}
	var c = [0]int{}
	log.Println(a,b,c)
	log.Printf("%p", &a)
	log.Printf("%p", &b)
	log.Printf("%p", &c)


}

func main(){
	test()
}
